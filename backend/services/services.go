package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

type Post struct {
	ID       int
	Title    string
	Details  string
	Tag      string
	PostType string
}

const (
	host = "db"
	port = 5432
)

// var user User
var (
	user User
	ctx  context.Context
)
var db *sql.DB = Connect_DB()

func GetMessage() (string, error) {
	//test
	return "Hello World", nil
}
func GetHealthCheck() (string, error) {
	return "Go fiber is good!", nil
}

func GetAllPost() ([]Post, error) {
	sqlStatement := `SELECT * FROM post`
	//Query all rows in table post
	rows, err := db.Query(sqlStatement)

	//Create postList slice's for store post row form rows
	postList := make([]Post, 0)

	// release connection resource when finish this function
	defer rows.Close()
	//loop for scan and push row to slice for return to API
	for rows.Next() {
		var onePost Post
		if err := rows.Scan(&onePost.ID, &onePost.Title, &onePost.Details,
			&onePost.Tag, &onePost.PostType); err != nil {
			log.Fatal(err)
		}
		fmt.Println(onePost)
		postList = append(postList, onePost)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(postList)
	return postList, nil
}
func GetAllUser() ([]User, error) {
	sqlStatement := "SELECT * FROM users;"
	//Query all rows in table users
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	// release connection resource when finish this function
	defer rows.Close()
	//Create list slice's for store post row form rows
	list := make([]User, 0)
	//loop for scan and push row to slice for return to API
	for rows.Next() {
		var oneUser User
		if err := rows.Scan(&oneUser.ID, &oneUser.Age, &oneUser.FirstName, &oneUser.LastName, &oneUser.Email); err != nil {
			log.Fatal(err)
		}
		list = append(list, oneUser)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return list, nil
}

func GetPostById(id string) Post {
	// pase id datatype string to int
	post_id, _ := strconv.Atoi(id)
	sqlStatement := `SELECT * FROM post Where id=$1`
	var post Post
	//Query one row from DB
	row := db.QueryRow(sqlStatement, post_id)
	err := row.Scan(&post.ID, &post.Title, &post.Details,
		&post.Tag, &post.PostType)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(post)
	default:
		panic(err)
	}

	fmt.Println(post)
	return post
}

func Connect_DB() *sql.DB {
	//Create psql Info for connect your Postgres DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("dbname"))
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