package services

import (
	database "daeng-market/databases"
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

const (
	host = "db"
	port = 5432
)

var (
	User model.User
	Post model.Post
	ctx  context.Context
)
var db *sql.DB = Connect_DB()
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

var dbasd = database.GetDB()

func GetMessage() (string, error) {
	//test
	return "Hello WorDogld", nil
}
func GetHealthCheck() (string, error) {
	return "Go fiber is good!", nil
}

func GetAllPost() ([]model.Post, error) {
	sqlStatement := `SELECT * FROM post`
	//Query all rows in table post
	rows, err := dbasd.Query(sqlStatement)

	//Create postList slice's for store post row form rows
	postList := make([]model.Post, 0)

	// release connection resource when finish this function
	defer rows.Close()
	//loop for scan and push row to slice for return to API
	for rows.Next() {
		var onePost model.Post
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
// func GetPostById(id string) model.Post {
// 	// pase id datatype string to int
// 	post_id, _ := strconv.Atoi(id)
// 	sqlStatement := `SELECT * FROM post Where id=$1`
// 	var post model.Post
// 	//Query one row from DB
// 	row := db.QueryRow(sqlStatement, post_id)
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
func GetPostByTag(tag1, tag2 string) ([]model.Post, error) {
	sqlStatement := `SELECT * FROM post Where tag=$1 OR tag=$2`

	rows, err := db.Query(sqlStatement, tag1, tag2)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]model.Post, 0)

	for rows.Next() {
		var onePost model.Post
		if err := rows.Scan(&onePost.ID, &onePost.Title, &onePost.Details,
			&onePost.Tag, &onePost.PostType); err != nil {
			log.Fatal(err)
		}
		fmt.Println(onePost)
		list = append(list, onePost)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(list)
	return list, nil
}
func GetAllUser() ([]model.User, error) {
	sqlStatement := "SELECT * FROM users;"
	//Query all rows in table users
	rows, err := dbasd.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	// release connection resource when finish this function
	defer rows.Close()
	//Create list slice's for store post row form rows
	list := make([]model.User, 0)
	//loop for scan and push row to slice for return to API
	for rows.Next() {
		var oneUser model.User
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
	//Query one row from dbasd
	row := dbasd.QueryRow(sqlStatement, post_id)
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
