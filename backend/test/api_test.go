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
// func Test_Create_Post_StatSuccess(t *testing.T) {
// 	app := fiber.New()

// 	controllers.Route(app)

// 	resp, err := app.Test(httptest.NewRequest("POST", "/post/create", nil))
// 	utils.AssertEqual(t, nil, err, "app.Test")
// 	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
// }
func Test_GetAllPost(t *testing.T) {
	app := fiber.New()

	controllers.Route(app)

	res, err := app.Test(httptest.NewRequest("GET", "/post/all", nil))
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 200, res.StatusCode, "Status code")
}
func Test_GetPostById(t *testing.T) {
	app := fiber.New()

	controllers.Route(app)
	res, err := app.Test(httptest.NewRequest("GET", "/post/getbyid/2", nil))
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 200, res.StatusCode, "Status code")
}
func Test_GetPostByTag(t *testing.T) {
	app := fiber.New()

	controllers.Route(app)
	res, err := app.Test(httptest.NewRequest("GET", "/post/getbytag/6/7", nil))
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 200, res.StatusCode, "Status code")
}

// func Test_CreatePost(t *testing.T) {
// 	app := fiber.New()

// 	controllers.Route(app)
// 	res, err := app.Test(httptest.NewRequest("POST", "/post/create", nil))
// 	utils.AssertEqual(t, nil, err, "app.Test")
// 	utils.AssertEqual(t, 200, res.StatusCode, "Status code")
// }
