package controllers

import (
	model "daeng-market/models"
	services "daeng-market/services"

	"github.com/gofiber/fiber/v2"
)

// User object
// type Post struct {
// 	PostId        int       `json:"post_id"`
// 	ProductName   string    `json:"product_name"`
// 	PostDate      time.Time `json:"post_date"`
// 	ProductOption string    `json:"product_option"`
// 	Price         float64   `json:"price"`
// 	Amount        int       `json:"amount"`
// 	PinId         int       `json:"pin_id"`
// 	TagId         int       `json:"tag_id"`
// }

// Users collection
type PostList []model.Post

func Route(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		msg, _ := services.GetMessage()
		return c.Status(200).JSON(msg)
	})
	app.Get("/health", func(c *fiber.Ctx) error {
		msg, _ := services.GetHealthCheck()
		return c.Status(200).JSON(msg)
	})

	//make post entry point
	post_entry_point := app.Group("/post")
	PostRoute(post_entry_point)
	//make user entry point
	user_entry_point := app.Group("/user")
	UserRoute(user_entry_point)

}

// func PostHandler(w http.ResponseWriter, r *http.Request) {
// 	u := PostList{
// 		Post{
// 			PostId:        1,
// 			ProductName:   "อยากได้ค้อนครับ",
// 			PostDate:      time.Now(),
// 			ProductOption: "ต้องการ",
// 			Price:         123,
// 			Amount:        1,
// 			PinId:         1234,
// 			TagId:         1346,
// 		},
// 		// User{
// 		// 	Firstname: "Jane",
// 		// 	Lastname:  "Wongsuwanjun",
// 		// 	Title:     "Miss.",
// 		// },
// 	}
// 	w.WriteHeader(http.StatusBadRequest)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(u)
// }
