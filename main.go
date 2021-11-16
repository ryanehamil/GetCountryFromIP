package main

import "fmt"

var verbose bool = false
var detail bool
var ip string
var properties []string

func main() {
	parseFlags()

	if verbose {
		fmt.Println("Verbose mode on")
	}
	data, err := lookup(ip, properties)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := getProperties(data)

	fmt.Println(result)
}
