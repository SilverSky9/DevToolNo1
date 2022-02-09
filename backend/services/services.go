package services

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}
// type Post struct {
// 	ID       int
// 	Title    string
// 	Details  string
// 	Tag      string
// 	PostType string
// }

const (
	host = "db"
	port = 5432
)

var user User
var db *sql.DB = Connect_DB()

func GetMessage() (string, error) {
	//test
	return "Hello World", nil
}
func GetHealthCheck() (string, error) {
	return "Go fiber is good!", nil
}
// func GetAllPost() ([]Post, error) {
// 	sqlStatement := `SELECT * FROM post`
// 	var post Post
// 	rows, err := db.Query(sqlStatement)
// 	var postList []Post

// 	switch err {
// 	case sql.ErrNoRows:
// 		fmt.Println("No rows were returned!")
// 	case nil:
// 		fmt.Println(post)
// 	default:
// 		panic(err)
// 	}
// 	for rows.Next() {
// 		var onePost Post
// 		if err := rows.Scan(&post.ID, &post.Title, &post.Details,
// 			&post.Tag, &post.PostType); err != nil {
// 			return postList, err
// 		}
// 		fmt.Println(err)
// 		postList = append(postList, onePost)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return postList, err
// 	}
// 	fmt.Println(postList)
// 	return postList, nil
// }
func GetUser() ([]User, error) {
	// sqlStatement := "SELECT * FROM users"
	rows, err := db.Query("SELECT * FROM users")
	// err := rows.Scan(&user.ID, &user.Age, &user.FirstName,
	// 	&user.LastName, &user.Email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []User
	// switch err {
	// case sql.ErrNoRows:
	// 	fmt.Println("No rows were returned!")
	// case nil:
	// 	fmt.Println(user)
	// 	// userList = append(userList, user)
	// default:
	// 	panic(err)
	// }
	for rows.Next() {
		var oneUser User
		if err := rows.Scan(&oneUser.ID, &oneUser.Age, &oneUser.FirstName, &oneUser.LastName, &oneUser.Email); err != nil {
			return list, err

		}
		fmt.Println(oneUser)
		list = append(list, user)
		// fmt.Print(userList)
	}
	if err = rows.Err(); err != nil {
		return list, err
	}
	return list, nil
}
// func GetPostById(id string) Post {
// 	sqlStatement := `SELECT * FROM post`
// 	var post Post
// 	row := db.QueryRow(sqlStatement)
// 	err := row.Scan(&post.ID, &post.Title, &post.Details,
// 		&post.Tag, &post.PostType)
// 	switch err {
// 	case sql.ErrNoRows:
// 		fmt.Println("No rows were returned!")
// 	case nil:
// 		fmt.Println(post)
// 	default:
// 		panic(err)
// 	}

// 	fmt.Println(post)
// 	return post
// }

func Connect_DB() *sql.DB {
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

// func Disconnect_DB() {
// 	db.Close()
// 	fmt.Println("Disconnected!")
// }
