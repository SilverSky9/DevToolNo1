package services

import (
	model "daeng-market/models"
	"reflect"
	"testing"
)

func TestGetAllPost(t *testing.T) {
	tests := []struct {
		name    string
		want    []model.Post
		wantErr bool
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllPost()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllPost() = %v, want %v", got, tt.want)
			}
		})
	}
}
