package main

import (
	"flag"
	"fmt"
	"lookupip/src/ipapi"
	"strings"
)

var verbose bool = false
var detail bool
var ip string
var properties []string

func main() {
	parseFlags()

	if verbose {
		fmt.Println("Verbose mode on")
	}
	data, err := ipapi.Lookup(ip, properties)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := ipapi.GetProperties(data, properties, detail)

	fmt.Println(result)
}

func parseFlags() {
	_ip := flag.String("ip", "", "IP address to lookup")
	_props := flag.String("p", "Country", "Properties to retrieve")
	flag.BoolVar(&detail, "d", false, "Show Detail")
	_verbose := flag.Bool("v", false, "Verbose output")
	flag.Parse()
	_loose := flag.Args()

	if *_verbose {
		verbose = true
	}
	if *_ip == "" {
		if len(_loose) == 0 {
		} else {
			ip = _loose[0]
		}
	} else {
		ip = *_ip
	}
	if _props != nil {
		properties = strings.Split(*_props, ",")
	}
}
