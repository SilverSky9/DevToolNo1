package services

import (
	database "daeng-market/databases"
	model "daeng-market/models"
	"daeng-market/queries"
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
