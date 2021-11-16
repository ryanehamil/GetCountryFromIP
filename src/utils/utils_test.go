package utils

import (
	"testing"
)

func TestCheckValidIP(t *testing.T) {

	var tests = []struct {
		ip      string
		want    bool
		explain string
	}{
		{"127.0.0.1", true, "Valid IP"},
		{"127.O.O.1", false, "Invalid IP"},
	}

	for _, test := range tests {
		got := CheckValidIP(test.ip)
		if got != test.want {
			t.Errorf("CheckValidIP(%q) = %v, want %v", test.ip, got, test.want)
		}
	}
}
