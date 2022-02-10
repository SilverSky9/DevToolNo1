package model

// type User struct {
// 	Name string `json:"name", bson:"name"`
// }
type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}