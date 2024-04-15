// TODO
// (1) - Write a program which prompts the user to enter integers and stores the integers in a sorted slice.The program should be written as a loop.
// (3) - Before entering the loop, the program should create an empty integer slice of size (length) 3.
// (4) - During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
// (5) - The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
// (6) - The slice must grow in size to accommodate any number of integers which the user decides to enter.
// (7) - The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Defining a slice of size 3
	slice := make([]int, 3)
	fmt.Println("Initialized the slice: ", slice)
	// Loop to take input from user
	for {
		// Taking input from user
		var input string

		fmt.Println("Enter an integer or 'X' to exit: ")
		fmt.Scan(&input)

		// If user enters 'X' then exit the loop
		if strings.ToLower(input) == "x" {
			fmt.Println("Final version: ", slice)
			break
		} else {
			fmt.Println("You entered: ", input)
			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Unsuccessful string to int conversion!")
				continue
			}
			// Append the input to the slice
			slice = append(slice, num)
			// Sorting the slice
			sort.Ints(slice)
			// Printing the sorted slice
			fmt.Println("Slice: ", slice)
		}
	}
}
