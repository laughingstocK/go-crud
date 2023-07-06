package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// import

func init() {

}

func main() {
	// connect db
	// Connection parameters
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "admin"
	dbPass := "password"
	dbName := "gocrud"

	// Create connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Open the database connection
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	// Connection successful
	fmt.Println("Connected to the MariaDB database!")

	// Close the database connection when done
	defer db.Close()

	// make route

	e := echo.New()
	e.Use(middleware.CORS())

	e.Start(":8080")
}
