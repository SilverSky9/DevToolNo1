package services

import (
	model "daeng-market/models"
	"daeng-market/queries"
)

func CreateUser(user model.User) (int, error) {
	user_id, err := queries.CreateUserQueries(user)
	if err != nil {
		return -1, err
	}
	return user_id, nil

}
func GetAllUser() ([]model.User, error) {
	resp, err := queries.GetAllUserQueries()
	if err != nil {
		return []model.User{}, err
	}

	return resp, nil
}
