package services

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

//compare pin -> gen jwt token -> save cookies, c for use fiber cookies
func RegisterNewToken(pin string, post_id int, c *fiber.Ctx) (int, error) {
	//get pin_id from post table
	pin_id, err := GetPinByPostId(post_id)
	if err != nil {
		return 500, err
	}
	//get pin from db
	hashed_pin, err := GetPinById(pin_id)
	if err != nil {
		return 500, err
	}
	//compare pin btw input and db
	validate_pin := CompareHashedPin(pin, hashed_pin)
	if !validate_pin {
		return 400, errors.New("eiei")
	}
	fmt.Println(validate_pin, pin)

	token, err := generateJwtToken(post_id)

	SaveCookies(c, post_id, token)

	return 200, nil
}

//save cookies name as post_id and sub is jwt token
func SaveCookies(c *fiber.Ctx, post_id int, token string) {
	cookie := fiber.Cookie{
		Name:    "post_" + strconv.Itoa(post_id),
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	}

	c.Cookie(&cookie)
}
