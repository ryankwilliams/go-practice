package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func example1() {
	// Read yaml from a file

	var filename string = "servers.yml"

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to read %s. Error: %s", filename, err)
	}

	type Server struct {
		Url      string
		Username string
		Password string
	}

	servers := make(map[string]Server)

	yaml.Unmarshal(content, &servers)

	fmt.Printf("Server 1 url: %s\n", servers["server1"].Url)
	fmt.Printf("Server 2 url: %s", servers["server2"].Url)
}

func example2() {
	// Write content to a yaml file

	var outFile string = "example2.yml"

	content := map[string]map[string]string{
		"cars": {
			"toyota": "4Runner",
		},
	}

	ymlContent, err := yaml.Marshal(content)
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(outFile, ymlContent, 0644)
	os.Remove(outFile)
}

func main() {
	example1()
	example2()
}
