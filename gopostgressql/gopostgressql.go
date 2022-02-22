package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {

	// urlExample := "postgres://username:password@localhost:5432/database_name"

	fmt.Println("**********START********")

	// Connect to a database
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "1. Unable to connect to database: %v\n", err)
		fmt.Println("**********Error********")
		os.Exit(1)
	} else {
		fmt.Printf("1. Connected to the DB: true [" + os.Getenv("DATABASE_URL") + "] \n")
		fmt.Println()
	}

	// Create a sequence
	statement := "CREATE SEQUENCE product_id_seq START 1;"
	_, err = conn.Exec(context.Background(), statement)

	if err != nil {
		fmt.Fprintf(os.Stderr, "2. Sequence create statement: %v\n", err)
		fmt.Println()
	} else {
		fmt.Printf("2. Sequence create statement: true\n")
		fmt.Println()
	}

	// Create a table
	statement = "CREATE TABLE product(id SERIAL PRIMARY KEY,price DECIMAL(14,2) NOT NULL,name TEXT NOT NULL,description TEXT NOT NULL,image TEXT NOT NULL);"
	_, err = conn.Exec(context.Background(), statement)

	if err != nil {
		fmt.Fprintf(os.Stderr, "3. Table create statement: %v\n", err)
		fmt.Println()
	} else {
		fmt.Printf("3. Table create statement: true\n")
		fmt.Println()
	}

	// Insert a value
	statement = "INSERT INTO product VALUES (nextval('product_id_seq'), 29.99, 'Return of the Jedi', 'Episode 6, Luke has the final confrontation with his father!', 'images/Return.jpg');"
	_, err = conn.Exec(context.Background(), statement)

	if err != nil {
		fmt.Fprintf(os.Stderr, "4. Insert statement: %v\n", err)
		fmt.Println()
	} else {
		fmt.Printf("4. Insert statement: true\n")
		fmt.Println()
	}

	// Query a value
	var name string
	var price float64

	err = conn.QueryRow(context.Background(), "select name, price from product where name='Return of the Jedi'").Scan(&name, &price)
	if err != nil {
		fmt.Printf("Connected to the DB: true\n")
		fmt.Println()
		fmt.Fprintf(os.Stderr, "5 QueryRow failed: %v\n", err)
		fmt.Println("**********Error********")
		os.Exit(1)
	} else {
		fmt.Println("Return values of the Table: ", name, price)
		fmt.Println()
	}

	defer conn.Close(context.Background())
	fmt.Println("**********DONE********")
}
