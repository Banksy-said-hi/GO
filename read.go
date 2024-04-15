// TODO
// Write a program which reads information from a file and represents it in a slice of structs.
// Assume that there is a text file which contains a series of names.
// Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
// Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
// Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file.
// Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file
// Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file
// After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.

package main

import (
	"fmt"
	"os"
)

func main() {

	type name struct {
		fname string
		lname string
	}

	var filename string
	var file *os.File
	var err error

	for {
		fmt.Println("Please enter the name of the text file: ")
		fmt.Scan(&filename)

		file, err = os.Open(filename)
		if err != nil {
			fmt.Println("Not such a file exists!")
		} else {
			break
		}
	}

	defer file.Close()

	var names []name

	for {
		var n name
		_, err := fmt.Fscan(file, &n.fname, &n.lname)
		if err != nil {
			break
		}
		names = append(names, n)
	}

	for index, n := range names {
		fmt.Println(index, n.fname, n.lname)
	}
}
