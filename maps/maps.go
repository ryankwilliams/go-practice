package main

import (
	"fmt"
)

func example1() {
	// Create an empty map with key/value pair of string:string
	fmt.Println("-- Example 1 --")
	map1 := make(map[string]string)
	map1["toyota"] = "4Runner"
	fmt.Println(map1["toyota"])
	fmt.Println("-- Example 1 --")
}

func example2() {
	// Create a map with key/value pair of string:string
	// plus initialize it
	fmt.Println("-- Example 2 --")
	map2 := map[string]string{
		"toyota": "4Runner",
	}
	fmt.Println(map2["toyota"])
	fmt.Println("-- Example 2 --")
}

func example3() {
	// Create a map with key/value pair of string:list[string]
	fmt.Println("-- Example 3 --")
	map3 := map[string][]string{
		"toyota": {
			"4Runner",
			"Rav4",
		},
	}
	fmt.Println(map3["toyota"])
	for k, v := range map3 {
		fmt.Printf("Key: %s\n", k)
		fmt.Printf("Value: %s\n", v)
		for i := 0; i < len(v); i++ {
			fmt.Printf("%d. %s\n", i+1, v[i])
		}
	}
	fmt.Println("-- Example 3 --")
}

func example4() {
	// Create a map with key/value pair of string:map[string][string]
	fmt.Println("-- Example 4 --")
	map4 := map[string]map[string]string{
		"foo": {
			"bar": "foo",
		},
	}
	fmt.Println(map4["foo"]["bar"])
	fmt.Println("-- Example 4 --")

	for _, v := range map4 {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}
}

func example5() {
	// Create a map and check to see if a key exists in the map or not
	fmt.Println("-- Example 5 --")
	map5 := make(map[string]string)
	var key string = "foo"
	var message string = ""
	_, exist := map5[key]
	if exist {
		message = "in"
	} else {
		message = "not in"
	}
	fmt.Printf("Key: %s is %s the map!", key, message)
	fmt.Println("-- Example 5 --")
}

func example6() {
	// Create a map and delete a key
	fmt.Println("-- Example 6 --")
	map6 := map[string]int{"one": 1, "two": 2}
	fmt.Printf("Before: %v\n", map6)
	delete(map6, "two")
	fmt.Printf("After: %v\n", map6)
	fmt.Println("-- Example 6 --")
}

func main() {
	example1()
	example2()
	example3()
	example4()
	example5()
	example6()
}
