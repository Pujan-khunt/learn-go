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
source /path/to/sql/file
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
