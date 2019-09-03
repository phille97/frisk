package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/phille97/frisk/lib"
)

var (
	checkUrl   string
	statusCode string
)

func init() {
	flag.StringVar(&checkUrl, "url", "https://google.se", "The url to check")
	flag.StringVar(&statusCode, "status_code", "200", "Expected status code")
}

func main() {
	flag.Parse()

	fmt.Println("# Frisk - a http, https, tcp healthcheck library")

	url, err := url.Parse(checkUrl)
	if err != nil {
		panic(err)
	}

	health_checker := lib.HttpChecker{}

	var opts = make(map[string]interface{})
	opts["expect_status_code"] = statusCode

	health := health_checker.GetHealth(url, opts)
	fmt.Printf("State: %s\n", health.State)
	fmt.Printf("Reason (%s): %s\n", health.Reason.Severity, health.Reason.Description)

	if health.State != lib.UP {
		os.Exit(1)
	}

	os.Exit(0)
}
