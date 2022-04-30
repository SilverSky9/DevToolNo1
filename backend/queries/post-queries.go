package queries

import (
	database "daeng-market/databases"
	model "daeng-market/models"
	"fmt"
	"strconv"
	"time"

	"github.com/lib/pq"
)

var dbasd = database.GetDB()

func CreatePostQueries(post model.Post) error {
	entryStm := `INSERT INTO post `
	columnStm := `(pin_id, product_name, post_date, product_option, price, amount, tag_id)`
	valuesStm := `VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING post_id`
	statement := entryStm + columnStm + valuesStm

	var post_id int
	load_sea_time, _ := time.LoadLocation("Asia/Kolkata")
	format_sea_time := time.Now().UTC().In(load_sea_time)

	err := db.QueryRow(statement, post.PinId, post.ProductName, format_sea_time, post.ProductOption, post.Price, post.Amount, post.TagId).Scan(&post_id)
	if err != nil {
		return err
	}
	fmt.Println("Post ID :", post_id)
	return nil
}

func GetAllPostQueries() ([]model.Post, error) {
	sqlStatement := `SELECT * FROM post`
	//Query all rows in table post
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return []model.Post{}, err
	}
	//Create postList slice's for store post row form rows
	// postList := make([]model.Post, 0)
	var postList []model.Post
	// release connection resource when finish this function
	defer rows.Close()
	//loop for scan and push row to slice for return to API
	for rows.Next() {
		var post model.Post
		err := rows.Scan(&post.PostId, &post.PinId, &post.ProductName, &post.PostDate,
			&post.ProductOption, &post.Price, &post.Amount, &post.TagId)
		if err != nil {

			return []model.Post{}, err
		}
		postList = append(postList, post)
	}

	return postList, nil
}

func GetPostByIdQueries(post_id int) (model.Post, error) {
	sqlStatement := `SELECT * FROM post WHERE post_id = $1`
	var post model.Post
	//Query one row from dbasd
	row := db.QueryRow(sqlStatement, post_id)
	err := row.Scan(&post.PostId, &post.PinId, &post.ProductName, &post.PostDate,
		&post.ProductOption, &post.Price, &post.Amount, &post.TagId)
	if err != nil {
		return model.Post{}, err
	}

	return post, nil
}

func GetPostByTagQueries(tag_1 string, tag_2 string) ([]model.Post, error) {
	sqlStatement := `SELECT * FROM post WHERE tag_id = $1 OR tag_id = $2`

	tag1_int, _ := strconv.Atoi(tag_1)
	tag2_int, _ := strconv.Atoi(tag_2)

	rows, err := db.Query(sqlStatement, tag1_int, tag2_int)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]model.Post, 0)

	for rows.Next() {
		var onePost model.Post
		err := rows.Scan(&onePost.PostId, &onePost.PinId, &onePost.ProductName, &onePost.PostDate,
			&onePost.ProductOption, &onePost.Price, &onePost.Amount, &onePost.TagId)
		if err != nil {
			return []model.Post{}, err
		}

		list = append(list, onePost)
	}
	if err = rows.Err(); err != nil {
		return []model.Post{}, err
	}

	return list, nil
}

func GetPostByNameQueries(post_name string) ([]model.Post, error) {
	// where_stm := `'%` + `$1` + `%'`
	sqlStatement := `SELECT * FROM post WHERE LOWER(product_name) LIKE '%'||$1||'%'`

	rows, err := db.Query(sqlStatement, post_name)
	// rows, err := dbasd.Query(sqlStatement, tag1_int, tag2_int)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]model.Post, 0)

	for rows.Next() {
		var onePost model.Post
		err := rows.Scan(&onePost.PostId, &onePost.PinId, &onePost.ProductName, &onePost.PostDate,
			&onePost.ProductOption, &onePost.Price, &onePost.Amount, &onePost.TagId)
		if err != nil {
			return []model.Post{}, err
		}

		list = append(list, onePost)
	}
	if err = rows.Err(); err != nil {
		return []model.Post{}, err
	}

	return list, nil
}

func GetPostByMultiTagQueries(tags []int) ([]model.Post, error) {
	var post_list []model.Post

	sqlStatement := `SELECT * FROM post WHERE tag_id = any ($1)`
	rows, err := db.Query(sqlStatement, pq.Array(tags))
	if err != nil {
		return []model.Post{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var onePost model.Post
		err := rows.Scan(&onePost.PostId, &onePost.PinId, &onePost.ProductName, &onePost.PostDate,
			&onePost.ProductOption, &onePost.Price, &onePost.Amount, &onePost.TagId)
		if err != nil {
			return []model.Post{}, err
		}

		post_list = append(post_list, onePost)
	}
	if err = rows.Err(); err != nil {
		return []model.Post{}, err
	}

	return post_list, nil
}
