package ipapi

import (
	"testing"
)

// TestBuildURL
// Tests that the buildURL function returns a correct URL
func TestBuildURL(t *testing.T) {
	var tests = []struct {
		ip      string
		want    string
		explain string
	}{
		{"127.0.0.1", "http://ip-api.com/json/127.0.0.1", "Valid Return"},
	}

	for _, test := range tests {
		got := buildURL(test.ip)
		if got != test.want {
			t.Errorf("CheckValidIP(%q) = %v, want %v", test.ip, got, test.want)
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
		{"8.8.8.8", "United States", "Valid Return"},
	}

	for _, test := range tests {
		got, err := Lookup(test.ip, []string{"country"})
		if err != nil {
			t.Errorf("CheckValidIP(%q) = %v, want %v", test.ip, err, test.want)
		}
		if test.want != got.Country {
			t.Errorf("CheckValidIP(%q) = %v, want %v", test.ip, got, test.want)
		}

	}
}
