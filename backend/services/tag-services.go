package services

import (
	model "daeng-market/models"
	"daeng-market/queries"
)

func CreateTag(tag model.Tag) (int, error) {
	tag_id, err := queries.CreateTagQueries(tag)
	if err != nil {
		return -1, err
	}

	return tag_id, nil
}
