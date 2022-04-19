package services

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var secretKey = "covid18"

//generate jwt token
func generateJwtToken(post_id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = strconv.Itoa(post_id)
	claims["exp"] = time.Now().Add(time.Second * 24).Unix()
	t, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return t, nil

}

//get sub token
func ExtractToken(tokenStr string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims, nil
}

func verifyToken(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return false
	}

	return token.Valid
}

func VerifyTokenCookies(c *fiber.Ctx, post_id string) bool {
	token := c.Cookies("post_" + post_id)

	vrf_res := verifyToken(token)

	return vrf_res
}
