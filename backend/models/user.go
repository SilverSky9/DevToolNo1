package model

// type User struct {
// 	Name string `json:"name", bson:"name"`
// }
type User2 struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

type User struct {
	UserId      int    `json:"user_id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Contact     string `json:"contact"`
	PhoneNumber string `json:"phone_number"`
}
