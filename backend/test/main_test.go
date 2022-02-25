package test

import (
	"daeng-market/services"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}
func TestGetHealthCheck(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test", want: "Go fiber is good!", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := services.GetHealthCheck()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHealthCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetHealthCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
