package services

import (
	model "daeng-market/models"
	"daeng-market/queries"
)

func CreateTag(tag_name string) (int, error) {
	tag_id, err := queries.CreateTagQueries(tag_name)
	if err != nil {
		return -1, err
	}

	return tag_id, nil
}

func GetAllTag() ([]model.Tag, error) {
	resp, err := queries.GetAllTagQueries()
	if err != nil {
		return []model.Tag{}, err
	}

	return resp, nil
}
