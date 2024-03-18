package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=27"

func main() {
	result, _ := url.Parse(myurl)
	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Port: ", result.Port()) //would be nil for the above example, also please note that its a method.
	fmt.Println("Raw Query: ", result.RawQuery)

	// We have another method called Query which provides url.values, thus can values of params can be extracted by keys.
	queryParams := result.Query()
	fmt.Println("url value of key \"list\":", queryParams["list"])

	// We can also loop over each params
	for key, val := range queryParams {
		fmt.Printf("Param for key \"%s\" is %s\n", key, val)
	}

	// We can also construct url
	urlToConstruct := &url.URL{
		Scheme:  "https",
		Host:    "prajwalbharadwaj.com",
		Path:    "/portfolio",
		RawPath: "user=someone",
	}
	fmt.Println("Constructed URL: ", urlToConstruct.String())
}
