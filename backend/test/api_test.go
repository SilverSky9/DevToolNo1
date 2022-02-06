package test

import (
	"daeng-market/services"
	"testing"
)

func Test_api(t *testing.T) {
	got, _ := services.GetHealthCheck()
	want := "Go fiber is good!"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
