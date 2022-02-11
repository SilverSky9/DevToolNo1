package services

import (
	database "daeng-market/databases"
	model "daeng-market/models"
	"daeng-market/queries"
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

var dbasd = database.GetDB()

func CreatePost(post model.Post, tag model.Tag, user model.User, pin model.Pin) error {
	tag_id, err := CreateTag(tag)
	if err != nil {
		return err
	}
	post.TagId = tag_id

	user_id, err := CreateUser(user)
	if err != nil {
		return err
	}

	pin_id, err := CreatePin(pin.Pin, user_id)
	if err != nil {
		return err
	}
	post.PinId = pin_id

	err = queries.CreatePostQueries(post)
	if err != nil {
		return err
	}

	return nil
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
		if err := rows.Scan(&onePost.PostId, &onePost.ProductName, &onePost.PostDate,
			&onePost.ProductOption, &onePost.Price, &onePost.Amount, &onePost.PinId, &onePost, &onePost.TagId); err != nil {
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
func GetPostByTag(tag1, tag2 string) ([]model.Post, error) {
	sqlStatement := `SELECT * FROM post Where tag=$1 OR tag=$2`

	rows, err := dbasd.Query(sqlStatement, tag1, tag2)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]model.Post, 0)

	for rows.Next() {
		var onePost model.Post
		if err := rows.Scan(&onePost.PostId, &onePost.ProductName, &onePost.PostDate,
			&onePost.ProductOption, &onePost.Price, &onePost.Amount, &onePost.PinId, &onePost, &onePost.TagId); err != nil {
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

func GetPostById(id string) model.Post {
	// pase id datatype string to int
	post_id, _ := strconv.Atoi(id)
	sqlStatement := `SELECT * FROM post Where id=$1`
	var post model.Post
	//Query one row from dbasd
	row := dbasd.QueryRow(sqlStatement, post_id)
	err := row.Scan(&post.PostId, &post.ProductName, &post.PostDate,
		&post.ProductOption, &post.Price, &post.Amount, &post.PinId, &post, &post.TagId)
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
