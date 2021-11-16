package ipapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

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

// Build URL to query IP-API
// https://ip-api.com/docs/api:json
func buildURL(ip string) string {
	return "http://ip-api.com/json/" + ip
}

func Lookup(ip string, properties []string) (*IPAPI, error) {
	var data *IPAPI

	if !utils.CheckValidIP(ip) {
		return data, errors.New("please enter a valid IP address")
	}

	url := buildURL(ip)
	resp, err := http.Get(url)

	if err != nil {
		// DebugOut(Error, err.Error())
		return data, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		// DebugOut(Error, err.Error())
		return data, err
	}
	if data.Status == "fail" {
		// DebugOut(Error, data.Message)
		return data, errors.New(data.Message)
	}
	return data, nil
}

func GetProperties(data *IPAPI, properties []string, detail bool) string {
	var result string
	var output string = ""

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
