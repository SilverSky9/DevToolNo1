package queries

import (
	model "daeng-market/models"
	"fmt"
)

func CreateUserQueries(user model.User) (int, error) {
	entryStm := `INSERT INTO user1 `
	columnStm := `(name, address, contact, phone_number)`
	valuesStm := `VALUES ($1, $2, $3, $4) RETURNING user_id`
	statement := entryStm + columnStm + valuesStm

	var user_id int

	err := db.QueryRow(statement, user.Name, user.Address, user.Contact, user.PhoneNumber).Scan(&user_id)
	if err != nil {
		return -1, err
	}
	fmt.Println("User ID:", user_id)
	return user_id, nil
}

func GetAllUserQueries() ([]model.User, error) {
	sqlStatement := "SELECT * FROM users;"
	//Query all rows in table users
	rows, err := db.Query(sqlStatement)
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
			return []model.User{}, err
		}
		list = append(list, oneUser)
	}
	if err = rows.Err(); err != nil {
		return []model.User{}, err
	}
	return list, nil
}

func GetUserByPostIdQueries(post_id string) ([]model.User, error) {
	selectStatement := "u.user_id, u.name, u.address, u.contact, u.phone_number"
	sqlStatement := "SELECT " + selectStatement + " FROM user1 AS u JOIN interested_post AS i ON u.user_id = i.user_id WHERE i.post_id = $1"

	rows, err := db.Query(sqlStatement, post_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var all_user []model.User

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.UserId, &user.Name, &user.Address, &user.Contact, &user.PhoneNumber)
		if err != nil {
			return []model.User{}, err
		}
		all_user = append(all_user, user)
	}

	return all_user, nil
}
