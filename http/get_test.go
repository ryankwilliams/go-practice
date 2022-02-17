// Module demonstrating how to use http package for get method
// https://pkg.go.dev/net/http
package main

import (
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	var bodyText string = Get("http://example.com")
	if !strings.Contains(bodyText, "<!doctype html>") {
		t.Fatalf("API response does not contain the correct string!")
	}
}
