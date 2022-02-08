package services

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "db"
	port = 5432
)

type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

var db *sql.DB = Connect_DB()

func GetMessage() (string, error) {
	//test
	return "Hello World", nil
}
func GetHealthCheck() (string, error) {
	return "Go fiber is good!", nil
}
func CloseConnect_DB()  {
	db.Close()
}
func GetPost() User {
	sqlStatement := `SELECT * FROM users WHERE id=$1;`
	var user User
	row := db.QueryRow(sqlStatement, 3)
	err := row.Scan(&user.ID, &user.Age, &user.FirstName,
		&user.LastName, &user.Email)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(user)
	default:
		panic(err)
	}
	return user
}

func Connect_DB() *sql.DB {
	// connStr := "host=%s port=%d user=%s password=%s dbname=% sslmode=%s "
	// connStr = fmt.Sprintf(connStr, host, 5432, "Postgres", "magical", "DevTools_DB", "disable")
	connStr := "host=db port=5432 user=Postgres password=magical dbname=DevTools_DB sslmode=disable"
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("DevTools_DB"))

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}
