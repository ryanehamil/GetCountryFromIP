package utils

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

// Parse command line flags using go's flag package
//
// https://golang.org/pkg/flag/
func ParseFlags(ip *string, properties *string, detail *bool) {
	flag.StringVar(ip, "ip", "", "IP address to lookup")
	flag.StringVar(properties, "p", "", "Properties to retrieve")
	flag.BoolVar(detail, "d", false, "Show Detail")
	flag.Parse()
}

// For printing to the console
//
// Currently using fmt.Println
func PrintOut(msg ...interface{}) {
	fmt.Println(msg...)
}

// Check if a string is a valid ipv4 address
//
// Currently using regexp.MustCompile and regexp.MatchString to return boolean
func CheckValidIP(ip string) bool {
	re := regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	return re.MatchString(ip)
}

// Error handler for the program
//
// Checks if error !nil and prints the error to the console
func HandleError(err error) {
	if err != nil {
		PrintOut(err)
		Exit(1)
	}
}

// Simple Exit function to handle exit codes
//
// Might need to be changed to a more robust one
func Exit(code int) {
	if code != 0 {
		panic("exit")
	}
	os.Exit(code)
}
