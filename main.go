package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/laughingstocK/go-crud/author/repository"
	"github.com/laughingstocK/go-crud/author/usecase"

	_authorHttpDeliver "github.com/laughingstocK/go-crud/author/delivery/http"
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

	const (
		defaultName = "world"
	)

	var (
		addr = flag.String("addr", "localhost:50051", "the address to connect to")
	)

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

	flag.Parse()

	// Set up grpc a connection to the server.
	grpcConn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer grpcConn.Close()

	// make route

	e := echo.New()
	e.Use(middleware.CORS())

	mariadbAuthorRepo := repository.NewMariadbAuthorRepo(db)
	grpcAuthorRepo := repository.NewGrpcAuthorRepo(grpcConn)

	authorUsecase := usecase.NewAuthorUsecase(mariadbAuthorRepo, grpcAuthorRepo)
	_authorHttpDeliver.NewAuthorHandler(e, authorUsecase)

	e.Start(":8080")
}
