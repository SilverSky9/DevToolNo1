package services

import (
	model "daeng-market/models"
	"daeng-market/queries"
	"log"
)

func CreateUser(user model.User) (int, error) {
	user_id, err := queries.CreateUserQueries(user)
	if err != nil {
		return -1, err
	}

	return user_id, nil

}
func GetAllUser() ([]model.User, error) {
	sqlStatement := "SELECT * FROM users;"
	//Query all rows in table users
	rows, err := dbasd.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	// release connection resource when finish this function
	defer rows.Close()
	//Create list slice's for store post row form rows
	list := make([]model.User, 0)
	//loop for scan and push row to slice for return to API
	for rows.Next() {
		var oneUser model.User
		if err := rows.Scan(&oneUser.UserId, &oneUser.Name, &oneUser.Address, &oneUser.Contact, &oneUser.PhoneNumber); err != nil {
			log.Fatal(err)
		}
		list = append(list, oneUser)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return list, nil
}
