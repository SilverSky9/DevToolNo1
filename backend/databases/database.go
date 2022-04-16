package database

import (
	"context"
	model "daeng-market/models"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

func TestDB() {

}

const (
	host = "128.199.128.171"
	// host = "localhost"
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
		host, port, "postgres", "magical", "Market")
	// host, port, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open("postgres", psqlInfo)
	// db, err := sql.Open("postgres", "postgres:magical@tcp(127.0.0.1:5432)/Market?parseTime=true")
	// ^ This is request db.env to assign value
	if err != nil {
		// panic(err)
		fmt.Println(NotConnectDatabaseWarning())
	}
	err = db.Ping()
	if err != nil {
		// panic(err)
		fmt.Println(NotConnectDatabaseWarning())
	}
	return db
}

func GetDB() *sql.DB {
	return db
}
func NotConnectDatabaseWarning() error {
	return errors.New("Go Service isn't connect database, But if you want to test unitest is normal")
}
