// Write a program which prompts the user to first enter a name, and then enter an address.
// program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
// program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Defining a map
	m := make(map[string]string)

	// Taking input from user
	var name, address string
	fmt.Println("Enter a name: ")
	fmt.Scan(&name)
	fmt.Println("Enter an address: ")
	fmt.Scan(&address)

	// Adding name and address to the map
	m["name"] = name
	m["address"] = address

	// marshalling the map and save it in a byte slice
	jsonData, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Printing the map
	fmt.Println(string(jsonData))
}
