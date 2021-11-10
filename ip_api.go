package main

import (
	"encoding/json"
	"net/http"
)

// Build URL to query IP-API
// https://ip-api.com/docs/api:json
func buildURL(ip string) string {
	return "http://ip-api.com/json/" + ip
}

func Lookup(ip string, properties []string) map[string]interface{} {
	CheckValidIP(ip)
	url := buildURL(ip)
	resp, err := http.Get(url)

	if err != nil {
		debugOut(Error, err.Error())
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		debugOut(Error, err.Error())
	}
	if data["status"] == "fail" {
		debugOut(Error, data["message"].(string))
	}
	return data
}

func GetProperties(data map[string]interface{}) string {

	var output string = ""

	for _, property := range properties {
		found := false
		for key, value := range data {
			if key == property {
				if output != "" {
					output += ","
				}
				output += value.(string)
				found = true
			}
		}
		if !found {
			debugOut(Error, "Property '"+property+"' not found")
		}
	}

	return output
}
