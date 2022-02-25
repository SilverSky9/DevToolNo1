package services

import (
	"testing"
)

func TestCompareHashedPin(t *testing.T) {
	var testCase1_123 = "$2a$12$79KGrIyD3ZbUgSxF01pzT.XPLKqYjhHgk4V4Snq.kbmdqHxr.nuvS"
	var testCase2_9999 = "$2a$12$OseT21ElAf.QQustOal6J.u64am76QKch.kDHWJhTGQKlGiNUeGzy"
	type args struct {
		pin       string
		hashedPin string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//Test-Case
		{name: "test-case-compare-123", args: args{pin: "123", hashedPin: testCase1_123}, want: true},
		{name: "test-case-compare-9999", args: args{pin: "9999", hashedPin: testCase2_9999}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareHashedPin(tt.args.pin, tt.args.hashedPin); got != tt.want {
				t.Errorf("CompareHashedPin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashingPin(t *testing.T) {
	var testCase1_7834 = "7834"
	var testCase2_8912 = "8912"
	type args struct {
		pin string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		//Test-case
		{name: "test-case-hashingPin-7834", args: args{"7834"}, want: testCase1_7834, wantErr: false},
		{name: "test-case-hashingPin-8912", args: args{"8912"}, want: testCase2_8912, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := HashingPin(tt.args.pin)
			if CompareHashedPin(tt.want, got) != true {
				t.Errorf("HashingPin() = %v, want %v", got, tt.want)
			}
		})
	}
}
