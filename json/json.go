package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func example1() {
	// Read JSON file
	fmt.Println("-- Example 1 --")
	type Example1 struct {
		CarManufacture string
	}

	content, err := os.ReadFile("example1.json")
	if err != nil {
		log.Fatalf("Unable to read json file: %s", err)
	}

	jsonContent := Example1{}
	_ = json.Unmarshal(content, &jsonContent)
	fmt.Println(jsonContent)
	fmt.Printf("Car manufacture: %s", jsonContent.CarManufacture)
	fmt.Println("-- Example 1 --")
}

func example2() {
	// Read JSON file
	fmt.Println("-- Example 2 --")
	type Car struct {
		Make  string
		Model string
		Trim  string
		Color string
		Year  int
	}

	type Cars struct {
		Cars []Car
	}

	content, err := os.ReadFile("example2.json")
	if err != nil {
		log.Fatalf("Unable to read json file: %s", err)
	}

	cars := Cars{}
	_ = json.Unmarshal(content, &cars)

	for i := 0; i < len(cars.Cars); i++ {
		var car Car = cars.Cars[i]
		fmt.Printf("Make: %s, model: %s, year: %d, trim: %s, color: %s\n",
			car.Make, car.Model, car.Year, car.Trim, car.Color)
	}

	fmt.Println("-- Example 2 --")
}

func example3() {
	// fmt.Println("-- Example 3 --")
	// data := []string{"make", "toyota"}
	// d, _ := json.Marshal(data)
	// fmt.Println(string(d))
	// fmt.Println("-- Example 3 --")

	// Convert map to JSON and write to file
	fmt.Println("-- Example 3 --")
	myMap := map[string]string{"make": "toyota"}
	d, _ := json.MarshalIndent(myMap, "", "  ")
	fmt.Println(string(d))
	os.WriteFile("example3-output.json", d, 0644)

	type Example3 struct {
		Make string
	}

	content, err := os.ReadFile("example3-output.json")
	if err != nil {
		log.Fatalf("Unable to read json file: %s", err)
	}
	data := Example3{}
	json.Unmarshal(content, &data)
	fmt.Println(data.Make)
	fmt.Println("-- Example 3 --")
}

func main() {
	example1()
	example2()
	example3()
}
