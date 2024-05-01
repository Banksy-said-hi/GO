# GO
Golang Learning Process
- April 6th: Concurrency, Parallelism, Structs instead of classes, No ( Inheritance, Generic, and Constructors )

- April 8th: Pointer, Basic Data Types, Variable Scope, Allocation and Dealloaction of Memory, Garbage Collection

- April 9th: Garbage Collection, Comments, Printing, Integers, Ints, Floats, Strings, String Packages, Constants, Control Flow, Switch, Scan

- April 10th: & (memory address, pointer creation), * ( value of the memory pointer ), % ( remainder ):
    -     var b *int = &a // b holds the memory address of a
 
- April 12th: Protocols and formats,  Request for comments ( RFC ), HTML, URI, HTTP, Encode and decode protocol packages, JSON marshalling and unmarshalling

- April 13th: File access ( open, read, write, close, seek )

<img width="855" alt="Screenshot 2024-04-16 at 12 02 13 PM" src="https://github.com/Banksy-said-hi/GO/assets/72816123/94863a8e-2fcf-4849-a137-780ad20f6155">


=================================================================================================================================

**"Passing a slice passes a reference to the original underlying array, while passing an array passes a copy of the entire array."**

Exactly. When you pass a slice to a function, any modifications made to the elements of the slice within that function will affect the original slice outside of it. This is because a slice contains a reference to the underlying array, so changes made to the slice are reflected in the original array.

On the other hand, when you pass an array to a function, you're passing a copy of the entire array. Therefore, modifications made to the array within the function won't affect the original array outside of it. The function receives a separate copy of the array, so changes made to this copy don't impact the original array.

=================================================================================================================================

**Empty Interface**

An empty interface in Go can hold values of any type because it doesn't specify any methods. It's effectively the interface type that specifies zero methods. This means you can assign any value to a variable of type empty interface.

```
var i interface{} = 42
var j interface{} = "hello"
var k interface{} = 3.14
```

=================================================================================================================================

<img width="908" alt="Screenshot 2024-05-01 at 7 51 17 PM" src="https://github.com/Banksy-said-hi/GO/assets/72816123/704af8e1-89ba-42b1-b2ce-51454a22ad1a">

=================================================================================================================================

**Interleaved Execution**

Interleaved Execution Analogy: Cooking in a Shared Kitchen

Imagine a shared kitchen where multiple chefs are preparing different dishes simultaneously. Each chef has their own work station and ingredients.

Interleaved Execution:
As the chefs work on their respective dishes, they share access to common resources such as cooking utensils, stovetops, and counters.
While Chef A might be chopping vegetables, Chef B could be sautéing meat, and Chef C could be baking a dessert.
As one chef completes a step in their recipe, they temporarily pause to allow another chef to use the same resources.
This interleaved process continues, with chefs taking turns using shared resources to make progress on their dishes. Overall, multiple dishes are prepared concurrently, maximizing kitchen efficiency.

=================================================================================================================================

**Goroutines**

Goroutines Analogy: Team Collaboration on a Project

Imagine a team of workers collaborating on a project, such as building a house. Each worker has their own specific tasks to complete, but they need to work together to finish the project efficiently.

Each worker represents a goroutine in the Go programming language.

Instead of one worker trying to complete the entire project alone, the work is divided into smaller tasks, and each worker (goroutine) focuses on a specific task.
For example, one worker might handle plumbing, another might focus on electrical wiring, and a third might work on carpentry.
The workers communicate and coordinate their efforts through channels, sharing information and collaborating on tasks that require input from multiple workers.
By dividing the work among multiple workers (goroutines) and allowing them to collaborate effectively, the project progresses efficiently, similar to how goroutines enable concurrent execution and cooperation in a Go program.




