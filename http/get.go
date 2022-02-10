// Module demonstrating how to use http package for get method
// https://pkg.go.dev/net/http
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodyText := string(body)

	log.Println(bodyText)
}
