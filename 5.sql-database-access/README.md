## Running MariaDB(MySQL) using Docker
```bash
docker run --name mysql-db \
-p 3306:3306 \
-e MYSQL_ROOT_PASSWORD="admin123" \
--rm -d mysql
```

## Accessing the DB
### 1. in Docker
```bash
docker exec -it mysql-db \ # Executes a command inside the running container with an interactive terminal
mysql -u root -p # Launches the mysql client, connecting as the root user and prompting for a password
```

### 2. In host machine
```bash
mysql -h 127.0.0.1 -P 3306 -u root -p
```

## Run the SQL file for creating table and inserting data
```bash
source /home/pujan/Desktop/learn/go/5.sql-database-access/data-access/create-tables.sql
```

## Methods and Functions in Go
### 1. Function 
- Example of a Normal function which takes 2 parameters (a, b) both of type integer.
- The function also returns an integer.
- This function can be executed by any .go file which imports the package containing this function since its exported.

```go
func Add(a int, b int) int {
    return a + b
}
```

### 2. Method
- A method is a function with a special receiver argument.
- The receiver binds the function to a specific type, making it a "method" of that type.
- Basically it relates to the OOP style, where you create a Class(in Java) and create state variables (fields) 
and also methods which set, update and modify those state variables.
```go
type User struct {
    Name string
    Email string
    IsActive bool
}

// Uses a pointer to receive the reference to the struct instance that called this method.
// Modifies the value of the struct instance.
func (u *User) Deactivate() {
    u.IsActive = false
}

func (u User) Greet() {
    fmt.Printf("Hello, my name is %s\n", u.Name)
}

func main() {
    user_1 := User{Name: "Pujan", Email: "pujankhunt2412@gmail.com", IsActive: true}
    user_1.Greet()

    fmt.Printf("Before: %s\n", user_1.IsActive) // Outputs True
    user_1.Deactivate()
    fmt.Printf("After: %s\n", user_1.IsActive) // Outputs False
}
```

## Misc

### 1. Go Get command
```bash
go get .
```
- Installs the dependencies required in the codebase to the project/go.mod file.

## Difference between sql.Open and db.Ping
The fundamental difference being that sql.Open doesn't actually open an connection to the DB.
It only prepares the database handle; it doesn't connect.

When you call sql.Open("mysql", dsn), it does the following tasks:
1. It checks of the mysql driver, if its registered and available.
2. formats the dsn to ensure a proper connection string.
3. allocatd and returns a pointer to the sql.DB struct, which is later used in the program to **actually open the connections**.

sql.Open is designed for lazy initialization and connections are only created when needed.

Your db can be offline, or your credientials could be wrong and the sql.Open still won't throw an error (will return nil).

---

The db.Ping() command does the following tasks:
1. It goes to the connection pool managed by the sql.DB instance
2. It tries to get a connection. If no idle connections exists, it will try to create a new connection to the DB,
using the connection string and the driver specified.
3. It sends a simple ping message to the db and waits for a reply.
4. It returns the connection to the pool.

If any of these steps fail, like the connection string is incorrect, db is offline, it will return an error.
This is why its a standard way to verify a connection to your database configuration at application startup.

## Understanding sql.DB struct
> sql.DB doesn't represent a single Connection, but rather a Connection Pool.

The sql.DB object is designed to be a long living object that you create once and share throughout your application.

Its safe for concurrent use from multiple goroutines.

Its main job is to manage the pool of underlying db connections to improve performance and resource management.

Opening a new database connection is an expensive operation.

### Why is opening a new db connection expensive?
Well you might think that as you need to make network requests to the DB server which might be located in another country, it might take a while for the request to reach there. But its not about the distance.

Its about the procedure that needs to be followed by the Server (Golang backend) and the DB (database server) to **establish a trusted communication channel**.
 
### What happens when you want to create a new db connection?
1. **TCP Handshake**: Before any data can be sent, you need to establish a reliable network connection. Which begins with a TCP handshake where both sides send each others' ACK numbers and synchronize using the receiving ACK numbers on both sides.

This takes atleast one full round-trip between your backend and the DB to only just agree to talk.

2. **SSL/TLS Negotiation**: For a secure connection both db and server need to exchange the certificates to verify the identity of other, agree on an encryption algorithm, and generate unique session keys, which involves public-key cryptograph, which is computationally intensive and will require several more round trips.

3. **Database Authentication**: The backend sends the DB credentials which will then be verified the running DB process on the DB server.

4. **Session Setup**: The DB server allocates memory and resources for this new session. It also setups information like character set, time zone, transaction isolation level, and prepares its internal state to handle queries from the backend.

### Reusing Connections To Save Time
> Reusing connections is possible and is the entire goal of connection management.

### What does it mean to reuse an existing db connection?
It means once an existing db connection has done all the procedure of establishing a secured communication channel, you reuse the same channel and pass the subsequent queries over the same channel, essentially bypassing the TCP Handshakes, SSL/TLS Negotiations and the DB auth setup.

Now that the application (Golang backend) can reuse existing applications, we can have multiple goroutines use the same pool and allow them to use the existing DB connections, **but not at the exact same time**. This is where the connection pool comes in.

Now that we have allowed multiple goroutines to use existing DB connections, we need to make sure that there are no 2 goroutines that are using the same DB connection to talk to the DB, since this will cause issues.

### Who will manage this conflict issue?
> Ans: a `Connection Pool`

A connection pool is a cache of active, ready-to-use db connections that are managed by the application's database driver. (in Go, this is the sql.DB object)

Hence sql.DB is about managing the connection pool.

If a goroutine needs to query the db, it needs to have *sql.DB object, which will provide the goroutine with a idle-connection (connection which isn't being used but is active and ready-to-use)

If a goroutine needs to query the db, it needs to have *sql.DB object, which will provide the goroutine with a idle-connection (connection which isn't being used but is active and ready-to-use).

The sql.DB also has a mutex which is used for locking. It prevents other goroutines to access DB connections which aren't idle and are being used by other goroutines.

Once the db query part is finished and the goroutine no longer requires the db connection, it will mark the connection as idle using 
```go
defer rows.Close()
```

Now the connection pool will add the connection to the list of idle connections ready to use by other goroutines.


## 1. Controllers (or Handlers)

Responsible for the following tasks:

- Receiving the HTTP request
- parsing the request body, headers, query string, params etc.
- calling the appropriate **Service** to do the actual work.
- Taking the result from the service and formatting a HTTP response with it.

What it shouldn't do:

- It shouldn't contain any business logic
- It shouldn't directly talk to the database server.


## 2. Service

Responsible for the following tasks:

- Using one or more **Repositories** to perform any kind of operation on data.
- Using business logic to perform complex validation. Eg. An artist cannot publish more than 10 albums in a year.
- Enriching data (e.g. after fetching the album from the repository, also fetch the artist details using another service.)

What it shouldn't do:

- It should have no knowledge of the HTTP request and response objects.

## 3. Repository

Responsible for the following tasks:

- Communicating with data sources (like database, file, or an external API)
- Abstracting the data persistence logic. It provides simple methods like GetById, Save, Update, Delete etc.
- Contains the actual database queries.

What it shouldn't do:

- It shouldn't contain any business logic. Its only concern is data in and out of storage.
