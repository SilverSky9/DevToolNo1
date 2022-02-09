package services

import (
	"database/sql"
	"fmt"
	"os"

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
func CloseConnect_DB() {
	db.Close()
}
func GetPost() ([]User, error) {
	sqlStatement := `SELECT * FROM users`
	var user User
	rows, err := db.Query(sqlStatement)
	// err := rows.Scan(&user.ID, &user.Age, &user.FirstName,
	// 	&user.LastName, &user.Email)
	var userList []User
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(user)
		// userList = append(userList, user)
	default:
		panic(err)
	}
	for rows.Next() {
		var oneUser User
		if err := rows.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email); err != nil {
			return userList, err
		}
		userList = append(userList, oneUser)
	}
	if err = rows.Err(); err != nil {
		return userList, err
	}
	return userList, nil
}

func Connect_DB() *sql.DB {
	// connStr := "host=%s port=%d user=%s password=%s dbname=% sslmode=%s "
	// connStr = fmt.Sprintf(connStr, host, 5432, "Postgres", "magical", "DevTools_DB", "disable")
	// connStr := "host=db port=5432 user=Postgres password=magical dbname=DevTools_DB sslmode=disable"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("dbname"))
	db, err := sql.Open("postgres", psqlInfo)

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
