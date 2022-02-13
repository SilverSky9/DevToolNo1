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
