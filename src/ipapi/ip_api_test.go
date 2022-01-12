package ipapi

import (
	"testing"

	"github.com/ryanehamil/lookupip/src/utils"
)

// TestBuildURL
// Tests that the buildURL function returns a correct URL
func TestBuildURL(t *testing.T) {
	var tests = []struct {
		ip      string
		want    string
		explain string
	}{
		{"127.0.0.1", "http://ip-api.com/json/127.0.0.1", "Valid IP"},
		{"  127.0.0.1 ", "http://ip-api.com/json/127.0.0.1", "Dirty IP"},
		{"127.foo.bar.1", "invalid IP address, failed CheckValidIP", "Unsantized input"},
	}

	for _, test := range tests {
		got, err := buildURL(test.ip)

		if err != nil {
			if err.Error() != test.want {
				t.Errorf("%q. error-buildURL(%q) = %v, want %v", test.explain, test.ip, got, test.want)
			}
		} else if got != test.want {
			t.Errorf("%q. buildURL(%q) = %v, want %v", test.explain, test.ip, got, test.want)
		}
	}
}

// TestLookup
// Tests that the lookup function returns IPAPI data and no error
func TestLookup(t *testing.T) {
	var tests = []struct {
		ip      string
		want    string
		explain string
	}{
		{"8.8.8.8", "United States", "Google Public DNS"},
		{"199.211.133.90", "United States", "Naval Oceanography"},
		{"116.202.3.251", "Germany", "JAM Software Germany"},
		{"", "Any IP", "Get my IP"},
	}

	for _, test := range tests {
		data, err := Lookup(test.ip)
		got := data.Country
		if err != nil && err.Error() != test.want {
			t.Errorf("%q. error-buildURL(%q) = %v, want %v", test.explain, test.ip, got, test.want)
		} else if test.want == "Any IP" {
			if !utils.CheckValidIP(data.Query) {
				t.Errorf("%q. error-buildURL(%q) = %v, want %v", test.explain, test.ip, got, test.want)
			}
		} else if got != test.want {
			t.Errorf("%q. buildURL(%q) = %v, want %v", test.explain, test.ip, got, test.want)
		}

	}
}

func TestGetProperties(t *testing.T) {
	var tests = []struct {
		properties string
		want       string
		explain    string
	}{
		{"", "United States", "No properties requested"},
		{"Country", "United States", "Properties: Country"},
		{"Country,ISP", "United States,Google LLC", "Properties: Country,ISP"},
	}

	for _, test := range tests {
		// This test relies on the lookup function
		ip := "8.8.8.8"
		// Use the IP-API to lookup anything
		data, _ := Lookup(ip)

		got := GetProperties(data, test.properties, false)
		if got != test.want {
			t.Errorf("%q. GetProperties(%q) = %v, want %v", test.explain, test.properties, got, test.want)
		}
	}
}
