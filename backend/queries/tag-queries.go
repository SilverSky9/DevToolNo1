package queries

import (
	model "daeng-market/models"
	"fmt"
)

func CreateTagQueries(tag model.Tag) (int, error) {
	entryStm := `INSERT INTO tag `
	columnStm := `(tag_name)`
	valuesStm := `VALUES ($1) RETURNING tag_id`
	statement := entryStm + columnStm + valuesStm

	var tag_id int

	err := db.QueryRow(statement, tag.TagName).Scan(&tag_id)
	if err != nil {
		return -1, err
	}

	fmt.Println("Tag ID:", tag_id)
	return tag_id, nil
}
