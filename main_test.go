package main

import (
	"testing"

	"github.com/ryanehamil/lookupip/src/ipapi"
	"github.com/ryanehamil/lookupip/src/utils"
)

// TestBuildURL
// Tests that the buildURL function returns a correct URL
func TestLookup(t *testing.T) {
	var tests = []struct {
		input   string
		want    string
		explain string
	}{
		{"", "Any IPV4", "Lookup my own IP"},
		{"8.8.8.8", "United States", "Lookup Google's DNS"},
	}

	for _, test := range tests {
		ip := test.input
		properties := ""
		// Use the IP-API to lookup anything
		data, _ := ipapi.Lookup(&ip, &properties)

		got := ipapi.GetProperties(data, properties, detail)

		if test.want == "Any IPV4" {
			if !utils.CheckValidIP(data.Query) {
				t.Errorf("%q. error-buildURL(%q) = %v, want %v", test.explain, test.input, got, test.want)
			}
		} else if got != test.want {
			t.Errorf("%q. lookup(%q) = %v, want %v", test.explain, test.input, got, test.want)
		}

	}
}
