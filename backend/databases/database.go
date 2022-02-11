package database

import (
	"context"
	model "daeng-market/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func TestDB() {

}

const (
	host = "localhost"
	port = 5432
)

// var user User
var (
	user model.User
	ctx  context.Context
)
var db *sql.DB = Connect_DB()

func init() {
	Connect_DB()
}

func Connect_DB() *sql.DB {
	//Create psql Info for connect your Postgres DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, "postgres", "1234", "Market")
	// host, port, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("dbname"))
	db, err := sql.Open("postgres", psqlInfo)
	// ^ This is request db.env to assign value
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func GetDB() *sql.DB {
	return db
}
