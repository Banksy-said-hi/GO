// TODO
// Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
// Each partition should be of approximately equal size.
// Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
// The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
// When sorting is complete, the main goroutine should print the entire sorted list.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Receive input from user a set of numbers seperated with a space
	fmt.Println("Enter a series of integers separated by spaces:")
	// Read the input from the user
	reader := bufio.NewReader(os.Stdin)
	// Read the input until a new line is encountered
	line, _ := reader.ReadString('\n')
	// Split the input into a list of integers
	var input []int
	// Loop through the input and convert each field to an integer
	for _, field := range strings.Fields(line) {
		// Convert the field to an integer
		num, err := strconv.Atoi(field)
		// Check if there is an error
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// Append the number to the input list
		input = append(input, num)
	}
	// Print the input list
	fmt.Println("Input List:", input)
	// Calculate the length of the input list
	length := len(input)
	// Calculate the size of each partition
	partitionSize := length / 4
	// Create a channel to communicate between the main goroutine and the sub goroutines
	ch := make(chan []int)
	// Create a goroutine to sort the first partition
	go sortPartition(input[:partitionSize], ch)
	// Create a goroutine to sort the second partition
	go sortPartition(input[partitionSize:2*partitionSize], ch)
	// Create a goroutine to sort the third partition
	go sortPartition(input[2*partitionSize:3*partitionSize], ch)
	// Create a goroutine to sort the fourth partition
	go sortPartition(input[3*partitionSize:], ch)
	// Create a list to store the sorted partitions
	var partitions [][]int
	// Loop through the number of partitions
	for i := 0; i < 4; i++ {
		// Receive the sorted partition from the channel
		partition := <-ch
		// Append the partition to the list of partitions
		partitions = append(partitions, partition)
	}
	// Merge the sorted partitions
	sorted := merge(partitions)
	// Print the sorted list
	fmt.Println("==============================")
	fmt.Println("Sorted List:", sorted)
}

// Function to sort a partition of an array
func sortPartition(partition []int, ch chan []int) {
	// Print the partition that will be sorted
	fmt.Println("Sorting:", partition)
	// Sort the partition
	for i := 0; i < len(partition); i++ {
		for j := i + 1; j < len(partition); j++ {
			if partition[i] > partition[j] {
				partition[i], partition[j] = partition[j], partition[i]
			}
		}
	}
	// Send the sorted partition to the channel
	ch <- partition
}

// Function to merge sorted partitions
func merge(partitions [][]int) []int {
	// Create a list to store the merged partitions
	merged := make([]int, 0)
	// Loop through the partitions
	for len(partitions) > 0 {
		// Find the smallest element in the partitions
		smallest := partitions[0][0]
		index := 0
		for i, partition := range partitions {
			if partition[0] < smallest {
				smallest = partition[0]
				index = i
			}
		}
		// Append the smallest element to the merged list
		merged = append(merged, smallest)
		// Remove the smallest element from the partition
		partitions[index] = partitions[index][1:]
		// Remove the partition if it is empty
		if len(partitions[index]) == 0 {
			partitions = append(partitions[:index], partitions[index+1:]...)
		}
	}
	// Return the merged list
	return merged
}
