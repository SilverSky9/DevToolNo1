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
	type args struct {
		pin string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test-case-hashingPin-4578", args: args{}, want: }
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashingPin(tt.args.pin)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashingPin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HashingPin() = %v, want %v", got, tt.want)
			}
		})
	}
}
