# Connect to postgresSQL db

The example is related to [pgx - PostgreSQL Driver and Toolkit](https://github.com/jackc/pgx)


### Step 1: Git clone

```sh
git clone https://github.com/thomassuedbroecker/go-access-postgres-example.git
cd go-access-postgres-example
```

### Step 2: Create a mod file (that file exists)

```sh
cd gopostgressql
go mod init example/gopostgressql
```

* Example output:

```sh
go: creating new go.mod: module example/gopostgressql
```

### Step 2: Create a go file (that file exists)

```sh
touch gopostgressql.go
```

### Step 3: Copy the code into the created file

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
	   fmt.Fprintf("Connected to the DB: true [" + os.Getenv("DATABASE_URL") + "] \n")
	}
	defer conn.Close(context.Background())

	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
	        fmt.Fprintf(os.Stderr, "Connected to the DB: true\n")
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)
}
```

### Step 4: Import the needed packages

```sh
go get github.com/jackc/pgx/v4
```

### Step 5: Set the enviornment variable

```sh
export DATABASE_URL="postgres://username:password@localhost:5432/database_name"
```

### Step 5: Execute the go program

```sh
go run  .
```

