/*Write a program which prompts the user to first enter a name, and then enter an address. 
Your program should create a map and add the name and address to the map using the keys 
“name” and “address”, respectively. Your program should use Marshal() to create a JSON object 
from the map, and then your program should print the JSON object.*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name string
	var address string
	var myMap2 map[string]string

	// prompt user for a name
	fmt.Printf("Please enter a name: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		name = scanner.Text()
	}

	// prompt user for an address
	fmt.Printf("\nPlease enter an address: ")

	scanner = bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		address = scanner.Text()
	}

	// create a map, add the keys "name” and “address”
	myMap := map[string]string{
		"name":    name,
		"address": address,
	}

	// use Marshal() to create a JSON object from the map
	myJson, _ := json.Marshal(myMap)

	// Unmarshal the JSON object
	_ = json.Unmarshal(myJson, &myMap2)

	// print the JSON object.
	fmt.Println("\nPrinting the unmarshalled JSON object:")

	for key, val := range myMap2 {
		fmt.Println(key, val)
	}
}
