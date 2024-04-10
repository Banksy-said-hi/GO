package main

import "fmt"

var userInput float64

func main() {

	fmt.Println("Please enter a floating point number: ")

	fmt.Scan(&userInput)

	fmt.Println("You entered: ", userInput)

	fmt.Println("Truncating the number...")

	fmt.Println("Result: ", int(userInput))
}
