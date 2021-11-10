package main

import "fmt"

var verbose bool = false
var ip string
var properties []string

func main() {
	parseFlags()

	result := GetProperties(Lookup(ip, properties))

	fmt.Println(result)
}
