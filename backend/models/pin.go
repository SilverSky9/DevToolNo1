package model

type Pin struct {
	PinId  int    `json:"pin_id"`
	Pin    string `json:"pin"`
	UserId int    `json:"user_id"`
}
