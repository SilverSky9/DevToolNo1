package services

import (
	"context"
	model "daeng-market/models"
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

// var dbasd = database.GetDB()

func GetMessage() (string, error) {
	//test
	return "Hello WorDogld", nil
}
func GetHealthCheck() (string, error) {
	return "Go fiber is good!", nil
}
