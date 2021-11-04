package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	pageStatusCode := getPageStatusCode("https://www.lider.cl")
	fmt.Printf("The status code is %d", pageStatusCode)
}

func getPageStatusCode(url string) int {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Could not get request")
	}

	return resp.StatusCode
}
