# Base Image for Go Applications
Use the `golang` base image for `Dockerfile`s.

# Initially Copy Only Files Which Store Dependencies
Initially only copy the `go.mod` and the `go.sum` files into the container to leverage the `Docker layer cache` which speeds up build times.

Docker builds images in a series of layers. Each instruction in the Dockerfile creates a new layer. On consequent rebuilds,
docker checks if the files in the current layer are same as the previous layer and reuses the existing layer from its cache,
instead of re-executing the instructions of that layer.
However, the moment the files in a layer changes, docker needs to rebuild the cache and execute the instructions from scratch.

The `go.mod` and `go.sum` files change very infrequently, whereas the source code (`*.go`) files change very frequently.

Here docker copies these files in the container and caches this layer by creating a checksum of both of these files.
The next time the image is built, it will again calculate the checksum to verify is the file content has changed.

1. **Cache-Hit**: If it hasn't, it will reuse the existing cache layer and would save time copying the 2 files. Note that the speed difference
isn't because docker doesn't need to re-copy 2 individual files.
2. **Cache-Miss**: If it has, then it will rebuild the layer and create the new checksum to store in the cache.

```dockerfile
COPY go.mod go.sum ./
```

The decision to either rebuild this layer or use the cache version boiles down to what happened to the previous layer.
Firstly docker will check if the previous layer used its cache or not, if it hasn't then it doesn't make any sense for any subsequent layers to use theirs.

Secondly, docker will compare the RAW strings `go mod download` with its previous run which was also `go mod download` and hasn't changed and hence will
use its cache, since the previous layer used its cache and in this layer both previous and current iteration runs have the same run command it obviously
means that the output of this command would be same as it were in the previous iteration run and hence the cache layer is used. Bingo! This is where the
huge time savings are earned. Since `go mod download` installs all dependencies including the huge standard library that go offers it would save quite a
lot of time.

```dockerfile
RUN go mod download
```

The part where docker compares the RAW strings with the previous and current iteration runs is dependent on what kind of command was used.

- For commands like `RUN`, `CMD`, `ENTRYPOINT` etc.: Docker primarily looks at the command string itself. The assumption is that if the command string is
the same then it must produce the same output as it did in the previous iteration run if and only if the above parent layers all have a cache-hit.

- For commands like `COPY` and `ADD`: Docker understands these commands involve external files and uses the method of checksums. The checksum of the
previous and current iteration runs only match if the content of any single/multiple file(s) changes.

# CGO_ENABLED=0 Why?

`CGO` stands for C Go which is a foreign function interface that allows Go programs to call C code. Enabling this, introduces dependencies and complexities
that developers often want to avoid.

The primary reason to disable CGO is to create a truly statically linked binary which has truly 0 dependencies.

If a Go application uses or imports a standard library package written that has a C implementation (like `net` or `os/user`), the Go compiler will link
against the system's standard C library (`glibc` in Linux). This creates a **dynamically linked executable**. Now this binary has a dependency on the C library
of this OS it was compiled on and will only run on those OSes which contain the same standard C library that it was compiled with.

If you disable CGO using `CGO_ENABLED=0` this will tell the Go compiler to use the pure Go implementations of the standard library package and avoid
linking with any of the standard C libraries.

Also by creating a self-contained executable, we don't need to include the system libraries in our final docker image, which helps keep the image size small.

# Why run Go binary as non-root user?
This practice of creating non-root users to run Go binaries follows the **principle of least privilege**.
This principle dictates that the user executing a process should only have the least amount of permissions required to run the process.

By default, applications inside the Docker container run as the `root` user. While the `root` user inside the containerize machine has fewer permissions
than the `root` user of the host machine, it still has the full administrative power of the container.

## How can not using this principel affect security?
For example, a user can find the flaw in the application (like buffer overflow or remote code execution vulnerability), they can exploit it to take control
of the running process.

If that process is `root`, the attacker gains `root` access to the entire container, which allows them to exploit in the following ways:
1. **Modify the Container Filesystem**: The attacker can install malicious tools like crypto miners or network scanners, modify the application binary
or delete everything causing denial of service.
2. **Attempt Container Breakout**: The attacker having `root` access to the container is in a much stronger position to exploit a kernel vulnerability and
gain `root` level access of the host machine.
3. **Sniff Network Traffic**: A `root` process can capture traffic from other containers running on the same host machine and the same docker network.

The best practice is to just create a user with the minimum permissions to run the application.
