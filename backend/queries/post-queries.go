package queries

import (
	database "daeng-market/databases"
	model "daeng-market/models"
	"fmt"
	"time"
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
