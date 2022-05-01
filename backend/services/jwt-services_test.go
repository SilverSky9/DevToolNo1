package services

import (
	"fmt"
	"testing"
)

func TestVerifyToken(t *testing.T) {
	tests := []struct {
		name    string
		post_id int
		want    bool
	}{
		//Test-Case
		{name: "testCase1", post_id: 1, want: true},
		{name: "testCase2", post_id: 2, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, _ := generateJwtToken(tt.post_id)
			if got := verifyToken(token); got != tt.want {
				fmt.Println(token, got)
				t.Errorf("verifyToken() = %v, want %v", got, tt.want)
			}
			fmt.Println(token)
		})
	}

}
