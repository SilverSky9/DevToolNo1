package queries

import (
	model "daeng-market/models"
	"fmt"
)

func CreatePinQueries(pin string, user_id int) (int, error) {
	entryStm := `INSERT INTO pin `
	columnStm := `(pin, user_id) `
	valuesStm := `VALUES ($1, $2) RETURNING pin_id`
	statement := entryStm + columnStm + valuesStm

	var pin_id int
	err := db.QueryRow(statement, pin, user_id).Scan(&pin_id)
	if err != nil {
		fmt.Println("pin err", err)
		return -1, err
	}

	return pin_id, nil
}

func GetPinByIdQueries(pin_id int) (string, error) {
	sqlStatement := `SELECT * FROM pin WHERE pin_id = $1`

	var pin model.Pin
	row := db.QueryRow(sqlStatement, pin_id)
	err := row.Scan(&pin.UserId, &pin.PinId, &pin.Pin)
	if err != nil {
		return "", err
	}

	return pin.Pin, nil

}
