// TODO
// Write a Bubble Sort program in Go.
// The program should prompt the user to type in a sequence of up to 10 integers.
// The program should print the integers out on one line, in sorted order, from least to greatest.
// Use your favorite search tool to find a description of how the bubble sort algorithm works.
// Use your favorite search tool to find a description of how the bubble sort algorithm works.

package main // Declare the package name for this Go program

import "fmt" // Import the "fmt" package for formatted I/O functions like Println and Scanln

func main() { // Define the main function that will be executed when the program runs
	var counter int                                                 // Declare a variable 'n' of type int to hold the number of integers to sort
	fmt.Println("Enter the number of integers to sort (up to 10):") // Print a message to the user
	fmt.Scanln(&counter)                                            // Read the user's input from the command line into the variable 'n'

	if counter > 10 { // If the user has entered a number greater than 10...
		fmt.Println("Input exceeds limit of 10. Only the first 10 numbers will be considered.") // Print a message to the user
		counter = 10                                                                            // Set 'n' to 10
	}

	slice := make([]int, counter)              // Create a slice of integers with length 'n'
	fmt.Println("Enter the integers:")         // Print a message to the user
	fmt.Println("Hit Enter after each number") // Print a message to the user

	for i := 0; i < counter; i++ { // For each index from 0 to 'n'...
		fmt.Scanln(&slice[i]) // Read the user's input from the command line into the slice at index 'i'
	}

	fmt.Println("Unsorted slice:", slice) // Print the unsorted slice to the user
	bubbleSort(slice)                     // Call the bubbleSort function with the slice as an argument
}

func bubbleSort(slice []int) { // Define the bubbleSort function that takes a slice of integers as an argument
	n := len(slice)           // Get the length of the slice and store it in the variable 'n'
	fmt.Println("Sorting...") // Print a message to the user

	for i := 0; i < n; i++ { // For each index from 0 to 'n'...
		fmt.Println("=====================") // Print the current state of the slice to the user
		fmt.Println("Slice:", slice)         // Print the current state of the slice to the user
		fmt.Println("n=", n)                 // Print the current value of 'n' to the user
		fmt.Println("i=", i)                 // Print the current value of 'i' to the user
		for j := 0; j < n-i-1; j++ {         // For each index from 0 to 'n-i-1'...
			fmt.Println("j=", j)       // Print the current value of 'j' to the user
			if slice[j] > slice[j+1] { // If the element at index 'j' is greater than the element at index 'j+1'...
				swap(slice, j)               // Swap the elements at index 'j' and 'j+1'
				fmt.Println("Slice:", slice) // Print the current state of the slice to the user
			}
		}
		fmt.Println("Slice:", slice)         // Print the current state of the slice to the user
		fmt.Println("=====================") // Print the current state of the slice to the user
	}

	fmt.Println("Sorted slice:", slice) // Print the sorted slice to the user
}

func swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}
