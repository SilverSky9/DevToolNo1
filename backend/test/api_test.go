package test

import (
	database "daeng-market/databases"
	"daeng-market/services"
	"testing"
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
// func TestHTTPGetPost(t *testing.T) {
// 	t.Run("it should return httpCode 200", func(t *testing.T) {
// 		req, err := http.NewRequest(http.MethodGet, "http://localhost:3000/post", nil)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		resp := httptest.NewRecorder()
// 		handler := http.HandlerFunc(controllers.PostHandler)
// 		handler.ServeHTTP(resp, req)
// 		if status := resp.Code; status != http.StatusOK {
// 			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
// 		}
// 	})
// }
