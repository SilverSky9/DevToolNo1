package test

import (
	"daeng-market/controllers"
	"daeng-market/services"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	utils "github.com/gofiber/utils"
)

func Test_api_GetHealthChecking(t *testing.T) {
	got, _ := services.GetHealthCheck()
	want := "Go fiber is good!"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func Test_api_home(t *testing.T) {
	got, _ := services.GetMessage()
	want := "Hello World"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// func Test_api_Post(t *testing.T) {
// 	want := {
// 					PostId:        1,
// 					ProductName:   "อยากได้ค้อนครับ",
// 					PostDate:      time.Now(),
// 					ProductOption: "ต้องการ",
// 					Price:         123,
// 					Amount:        1,
// 					PinId:         1234,
// 					TagId:         1346,
// 				},
// }
func Test_Post_StatSuccess(t *testing.T) {
	app := fiber.New()

	controllers.Route(app)

	resp, err := app.Test(httptest.NewRequest("GET", "/post", nil))
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")

}
