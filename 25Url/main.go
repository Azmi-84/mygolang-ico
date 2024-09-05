package main

import (
	"fmt"
	"net/url"
)

const youtubeURL = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println("Welcome to handling URL in GO lang")
	fmt.Println(youtubeURL)

	// parsing url
	result, err := url.Parse(youtubeURL)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Hostname())
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.Scheme)
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Printf("Type of query params are: %T\n", qParams)

	fmt.Println(qParams["v"])

	for _, val := range qParams {
		fmt.Println("Param is: ", val)
	}

	partsOfURL := &url.URL{
		Scheme: "https",
		Host:   "pkg.go.dev",
		Path:   "/io",
	}
	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)
}
