package users_storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	DATABASE = "quickmd_demo"
	USER     = "postgres"
	PASSWORD = "sa123"
)

var (
	Client *sql.DB
)

func init() {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)

	var err error

	Client, err = sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("connection successful")
}
