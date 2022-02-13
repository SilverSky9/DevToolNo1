package services

import "testing"

func TestGetHealthCheck(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test", want: "test", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHealthCheck()
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
