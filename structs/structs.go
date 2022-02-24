package main

import "fmt"

func example1() {
	fmt.Println("-- Example 1 --")
	type Phone struct {
		Manufacture string
		Device      string
		Model       string
		Color       string
		Storage     int
	}

	iPhone := Phone{
		Manufacture: "Apple",
		Device:      "iPhone",
		Model:       "12 Pro",
		Color:       "Silver",
		Storage:     128,
	}

	fmt.Printf("Manufacture: %s, Device: %s, Model: %s,"+
		"Color: %s, Storage: %dGB", iPhone.Manufacture,
		iPhone.Device, iPhone.Model, iPhone.Color, iPhone.Storage)
	fmt.Println("\n-- Example 1 --")
}

type Fruit struct {
	Name  string
	Color string
}

func createFruit(name string, color string) *Fruit {
	fruit := Fruit{
		Name:  name,
		Color: color,
	}
	return &fruit
}

func example2() {
	apple := createFruit("Apple", "Red")
	fmt.Printf("Apple can be the color '%s'\n", apple.Color)
}

func main() {
	example1()
	example2()
}
