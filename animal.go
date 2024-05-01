// TODO
// Write a program which allows the user to create a set of animals and to get information about those animals.
// Each animal has a name and can be either a cow, bird, or snake.
// With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created.
// Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake.
// The following table contains the three types of animals and their associated data.
// Animal Food eaten Locomtion method Spoken sound
// cow grass walk moo
// bird worms fly peep
// snake mice slither hsss

// Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
// Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
// Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

// Each “newanimal” command must be a single line containing three strings.
// The first string is “newanimal”.
// The second string is an arbitrary string which will be the name of the new animal.
// The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
// Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

// Each “query” command must be a single line containing 3 strings.
// The first string is “query”.
// The second string is the name of the animal.
// The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
// Your program should process each query command by printing out the requested data.

// Define an interface type called Animal which describes the methods of an animal.
// Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
// The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
// Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
// When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.

package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

// Snake struct
type Cow struct {
	name string
}

// Snake struct
func (c Cow) Eat() {
	fmt.Println("grass")
}

// Snake struct
func (c Cow) Move() {
	fmt.Println("walk")
}

// Snake struct
func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird struct
type Bird struct {
	name string
}

// Bird struct
func (b Bird) Eat() {
	fmt.Println("worms")
}

// Bird struct
func (b Bird) Move() {
	fmt.Println("fly")
}

// Bird struct
func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake struct
type Snake struct {
	name string
}

// Snake struct
func (s Snake) Eat() {
	fmt.Println("mice")
}

// Snake struct
func (s Snake) Move() {
	fmt.Println("slither")
}

// Snake struct
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	// Create a map to store the animals
	animals := make(map[string]Animal)
	// Loop forever
	for {
		// Prompt the user
		fmt.Print("> ")
		// Read the command
		var command string
		fmt.Scan(&command)
		// Check if the command is other than "newnaimal" or "query"
		if command != "newanimal" && command != "query" {
			fmt.Println("Invalid command")
			continue
		} else if command == "newanimal" {
			// Read the name and the type of the animal
			var name, animalType string
			fmt.Scan(&name, &animalType)
			// Check if the animalType is among options
			if animalType != "cow" && animalType != "bird" && animalType != "snake" {
				fmt.Println("Invalid animal type")
				continue
			}
			// Create the animal
			switch animalType {
			case "cow":
				animals[name] = Cow{name}
			case "bird":
				animals[name] = Bird{name}
			case "snake":
				animals[name] = Snake{name}
			}
			fmt.Println("Created it!")
			// If the command is query
		} else if command == "query" {
			// Read the name and the info
			var name, info string
			fmt.Scan(&name, &info)
			// Check if the animal exists
			if _, ok := animals[name]; !ok {
				fmt.Println("Animal not found")
				continue
			}
			// Get the info
			switch info {
			case "eat":
				animals[name].Eat()
			case "move":
				animals[name].Move()
			case "speak":
				animals[name].Speak()
			}
		}
	}
}
