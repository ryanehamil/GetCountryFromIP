package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Verbosity int

const (
	Info Verbosity = iota
	Warning
	Error
	Debug
)

func debugOut(level Verbosity, text string) {
	switch level {
	case Info:
		if verbose {
			fmt.Println("[INFO] " + text)
		} else {
			fmt.Println(text)
		}
	case Warning:
		if verbose {
			fmt.Println("[WARNING] " + text)
		}
	case Error:
		fmt.Println("[ERROR] " + text)
		os.Exit(1)
	case Debug:
		if verbose {
			fmt.Println("[DEBUG] " + text)
		}
	}
}

func checkValidIP(ip string) bool {
	re := regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	return re.MatchString(ip)
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
			debugOut(Error, "Please enter an IP address or use -h for help")
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
