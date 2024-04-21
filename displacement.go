// TODO
// Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement.
// Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.
// s = ½ a t2 + vot + so
// Where s is the displacement, a is the acceleration, v0 is the initial velocity, t is the time, and so is the initial displacement.

// You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so.
// GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement
// The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

package main

import "fmt"

func main() {
	var acceleration, initialVelocity, initialDisplacement, time float64

	fmt.Println("Enter the value for acceleration:")
	fmt.Scanln(&acceleration)

	fmt.Println("Enter the value for initial velocity:")
	fmt.Scanln(&initialVelocity)

	fmt.Println("Enter the value for initial displacement:")
	fmt.Scanln(&initialDisplacement)

	fmt.Println("Enter the value for time:")
	fmt.Scanln(&time)

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Println("Displacement after", time, "seconds:", fn(time))
}

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return 0.5*acceleration*time*time + initialVelocity*time + initialDisplacement
	}
}
