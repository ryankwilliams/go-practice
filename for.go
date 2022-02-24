package main

import (
	"fmt"
	"log"
)

// TODO: Add unit test
func main() {
	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	map1 := map[string]string{"name": "ryan"}
	for key, value := range map1 {
		fmt.Println(key)
		fmt.Println(value)
	}

	arr := []string{"A", "B", "C"}
	fmt.Println(arr[0])

	map2 := map[string]map[string]string{"parentKey": {"childKey": "childValue"}}
	fmt.Println(map2["parentKey"]["childKey"])

	for key, value := range map2["parentKey"] {
		fmt.Printf("Key: %s, value: %s", key, value)
	}
}
