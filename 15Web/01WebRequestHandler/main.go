package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://go.dev/doc/tutorial/web-service-gin"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// fmt.Println("Response: ", response)
	defer response.Body.Close()
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	data := string(databytes)

	fmt.Println("Data: ", data)
}
