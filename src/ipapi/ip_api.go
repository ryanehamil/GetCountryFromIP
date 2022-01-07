package ipapi

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/ryanehamil/lookupip/src/utils"
)

// IP-API data struct
// https://ip-api.com/docs/api:json
type IPAPI struct {
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continentCode"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"countryCode"`
	Region        string  `json:"region"`
	RegionName    string  `json:"regionName"`
	City          string  `json:"city"`
	Zip           string  `json:"zip"`
	Lat           float32 `json:"lat"`
	Lon           float32 `json:"lon"`
	Timezone      string  `json:"timezone"`
	Offset        int     `json:"offset"`
	Currency      string  `json:"currency"`
	ISP           string  `json:"isp"`
	Org           string  `json:"org"`
	AS            string  `json:"as"`
	ASName        string  `json:"asname"`
	Reverse       string  `json:"reverse"`
	Hosting       bool    `json:"hosting"`
	Query         string  `json:"query"`
	Status        string  `json:"status"`
	Message       string  `json:"message"`
}

// Add the .String() method to IPAPI
func (i *IPAPI) String() string {
	return i.Query
}

// Build URL to query IP-API
// https://ip-api.com/docs/api:json
func buildURL(ip string) string {
	return "http://ip-api.com/json/" + ip
}

// Get IP-API data about IP
//
// https://ip-api.com/docs/api:json
func Lookup(ip string, properties string) *IPAPI {
	var data *IPAPI

	if ip == "" {
		if properties == "" {
			properties = "Query"
		}
	} else {
		if properties == "" {
			properties = "Country"
		}
		if !utils.CheckValidIP(ip) {
			utils.PrintOut("Invalid IP address during Lookup")
			utils.Exit(1)
		}

	}
	url := buildURL(ip)
	resp, err := http.Get(url)
	utils.HandleError(err)

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	utils.HandleError(err)

	if data.Status == "fail" {
		utils.PrintOut("IP-API returned an error: " + data.Message)
		utils.Exit(1)
	}
	return data
}

func GetProperties(data *IPAPI, properties_string string, detail bool) string {
	result := ""
	output := ""
	properties := strings.Split(properties_string, ",")

	for _, property := range properties {
		datafield := reflect.Indirect(reflect.ValueOf(data)).FieldByName(property).String()
		if datafield != "" {
			result = datafield
		} else {
			result = "Not found"
		}
		if detail {
			output += property + ": " + result + "\n"
		} else {
			if output != "" {
				output += ","
			}
			output += result
		}
	}
	return output
}
