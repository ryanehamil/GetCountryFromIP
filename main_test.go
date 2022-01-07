package main

import (
	"testing"

	"github.com/ryanehamil/lookupip/src/ipapi"
	"github.com/ryanehamil/lookupip/src/utils"
)

// TestBuildURL
// Tests that the buildURL function returns a correct URL
func TestBuildURL(t *testing.T) {
	var tests = []struct {
		input   string
		want    string
		explain string
	}{
		{"nothing", "any ip address", "Valid Return"},
	}

	for _, test := range tests {
		// Use the IP-API to lookup anything
		got := ipapi.Lookup(ip, properties).String()
		// use regex to check if got is a valid ipv4 address

		if !utils.CheckValidIP(got) {
			t.Errorf("ipapi.Lookup returned %s using input %s, want %s", got, test.input, test.want)
		}
	}
}
