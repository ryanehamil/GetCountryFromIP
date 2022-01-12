package main

import (
	"flag"

	"github.com/ryanehamil/lookupip/src/ipapi"
	"github.com/ryanehamil/lookupip/src/utils"
)

func main() {
	ip := flag.String("ip", "", "IP address to lookup")
	properties := flag.String("p", "", "Properties to retrieve")
	detail := flag.Bool("d", false, "Show Detail")
	flag.Parse()

	// Use the IP-API to lookup the IP address
	data, err := ipapi.Lookup(*ip)
	utils.HandleError(err)

	// Format the data to a string
	result := ipapi.GetProperties(data, *properties, *detail)

	// Print result with PrintOut
	utils.PrintOut(result)
	utils.Exit(0)
}
