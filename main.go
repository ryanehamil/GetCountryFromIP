package main

import (
	"github.com/ryanehamil/lookupip/src/ipapi"
	"github.com/ryanehamil/lookupip/src/utils"
)

var detail bool
var ip string
var properties string

func main() {
	// Parse command line flags
	utils.ParseFlags(&ip, &properties, &detail)

	// Use the IP-API to lookup the IP address
	data := ipapi.Lookup(ip)

	// Format the data to a string
	result := ipapi.GetProperties(data, properties, detail)

	// Print result with PrintOut
	utils.PrintOut(result)
	utils.Exit(0)

}
