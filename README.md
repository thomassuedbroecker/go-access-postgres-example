# Connect to postgresSQL db

The example is related to [pgx - PostgreSQL Driver and Toolkit](https://github.com/jackc/pgx)

### Step 1: Create a mod file

```sh
go mod init example/gopostgressql
```

* Example output:

```sh
go: creating new go.mod: module example/gopostgressql
```

### Step 2: Create a go file

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
	}
	defer conn.Close(context.Background())

	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
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


example/gopsql


Source https://github.com/jackc/pgx# go-access-postgres-example
