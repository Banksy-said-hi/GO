// Assignment:
// Write a program which prompts the user to enter a string.
// The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
// The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
// The program should print “Not Found!” otherwise.
// The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

// Examples: The program should print “Found!” for the following example entered strings,
// “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
// The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var userInput string

	fmt.Println("Please enter a string: ")

	fmt.Scan(&userInput)

	userInput = strings.ToLower(userInput)

	switch {
	case (userInput[0] == 'i') && (userInput[len(userInput)-1] == 'n') && (strings.Contains(userInput, "a")):
		fmt.Println("Yeahh! Found!")
	default:
		fmt.Println("Sorry! Not Found!")
	}
}
