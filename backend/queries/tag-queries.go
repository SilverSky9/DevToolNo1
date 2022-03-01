package queries

import (
	model "daeng-market/models"
	"fmt"
)

func CreateTagQueries(tag_name string) (int, error) {
	entryStm := `INSERT INTO tag `
	columnStm := `(tag_name)`
	valuesStm := `VALUES ($1) RETURNING tag_id`
	statement := entryStm + columnStm + valuesStm

	var tag_id int

	err := db.QueryRow(statement, tag_name).Scan(&tag_id)
	if err != nil {
		return -1, err
	}

	fmt.Println("Tag ID:", tag_id)
	return tag_id, nil
}

func GetAllTagQueries() ([]model.Tag, error) {
	sqlStatement := `SELECT * FROM tag`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return []model.Tag{}, err
	}

	var tagList []model.Tag

	defer rows.Close()

	for rows.Next() {
		var tag model.Tag
		err := rows.Scan(&tag.TagId, &tag.TagName)
		if err != nil {

			return []model.Tag{}, err
		}
		tagList = append(tagList, tag)
	}

	return tagList, nil
}

func GetTagByNameQueries(tag_name string) (model.Tag, error) {
	sqlStatement := `SELECT * FROM tag WHERE tag_name = $1`
	var tag model.Tag
	//Query one row from dbasd
	row := db.QueryRow(sqlStatement, tag_name)
	err := row.Scan(&tag.TagId, &tag.TagName)
	if err != nil {
		fmt.Println(err)
		return model.Tag{}, err
	}

	return tag, nil
}
