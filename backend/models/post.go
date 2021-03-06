package model

import "time"

type Post struct {
	PostId        int       `json:"post_id"`
	PinId         int       `json:"pin_id"`
	ProductName   string    `json:"product_name"`
	PostDate      time.Time `json:"post_date"`
	ProductOption string    `json:"product_option"`
	Price         float64   `json:"price"`
	Amount        int       `json:"amount"`
	TagId         int       `json:"tag_id"`
}
