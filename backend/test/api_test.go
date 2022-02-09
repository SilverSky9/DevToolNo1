package test

import (
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
