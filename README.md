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


==============================================================================================================================================================================================================

**"Passing a slice passes a reference to the original underlying array, while passing an array passes a copy of the entire array."**

Exactly. When you pass a slice to a function, any modifications made to the elements of the slice within that function will affect the original slice outside of it. This is because a slice contains a reference to the underlying array, so changes made to the slice are reflected in the original array.

On the other hand, when you pass an array to a function, you're passing a copy of the entire array. Therefore, modifications made to the array within the function won't affect the original array outside of it. The function receives a separate copy of the array, so changes made to this copy don't impact the original array.

==============================================================================================================================================================================================================

**Empty Interface**

An empty interface in Go can hold values of any type because it doesn't specify any methods. It's effectively the interface type that specifies zero methods. This means you can assign any value to a variable of type empty interface.

```
var i interface{} = 42
var j interface{} = "hello"
var k interface{} = 3.14
```

==============================================================================================================================================================================================================

<img width="908" alt="Screenshot 2024-05-01 at 7 51 17 PM" src="https://github.com/Banksy-said-hi/GO/assets/72816123/704af8e1-89ba-42b1-b2ce-51454a22ad1a">

==============================================================================================================================================================================================================

**Interleaved Execution**

Interleaved Execution Analogy: Cooking in a Shared Kitchen

Imagine a shared kitchen where multiple chefs are preparing different dishes simultaneously. Each chef has their own work station and ingredients.

Interleaved Execution:
As the chefs work on their respective dishes, they share access to common resources such as cooking utensils, stovetops, and counters.
While Chef A might be chopping vegetables, Chef B could be sautéing meat, and Chef C could be baking a dessert.
As one chef completes a step in their recipe, they temporarily pause to allow another chef to use the same resources.
This interleaved process continues, with chefs taking turns using shared resources to make progress on their dishes. Overall, multiple dishes are prepared concurrently, maximizing kitchen efficiency.

==============================================================================================================================================================================================================

**Goroutines**

Goroutines Analogy: Team Collaboration on a Project

Imagine a team of workers collaborating on a project, such as building a house. Each worker has their own specific tasks to complete, but they need to work together to finish the project efficiently.

Each worker represents a goroutine in the Go programming language.

Instead of one worker trying to complete the entire project alone, the work is divided into smaller tasks, and each worker (goroutine) focuses on a specific task.
For example, one worker might handle plumbing, another might focus on electrical wiring, and a third might work on carpentry.
The workers communicate and coordinate their efforts through channels, sharing information and collaborating on tasks that require input from multiple workers.
By dividing the work among multiple workers (goroutines) and allowing them to collaborate effectively, the project progresses efficiently, similar to how goroutines enable concurrent execution and cooperation in a Go program.

==============================================================================================================================================================================================================

**Concurrency** is often used to handle tasks that involve waiting for I/O operations (like reading from files or making network requests) or to improve responsiveness in applications by allowing them to perform other tasks while waiting.

**Parallelism** is particularly useful for computationally intensive tasks that can be divided into smaller independent parts, allowing them to be executed simultaneously on multiple CPU cores, thus reducing overall computation time.

==============================================================================================================================================================================================================

**Clock Cycles**
Clock cycles, also known as clock ticks, are the basic units of time used in computer processors to synchronize the execution of instructions. Each clock cycle represents a fixed period of time during which the processor performs various operations, such as fetching, decoding, executing, and writing back instructions.

**Clock Frequency**
Clock frequency, measured in hertz (Hz), represents the speed at which the processor's clock generates pulses. It indicates how many clock cycles the processor can execute per second. For example, a processor with a clock frequency of 3.0 GHz (gigahertz) can perform 3 billion clock cycles per second.

Let's imagine you have a toy car. Every time you push a button on the car, it moves a little bit. Now, let's pretend that button makes a "beep" sound every time you push it.

Clock Cycle: Each time you hear the "beep" sound, it means one small step has been completed. Maybe it's the car moving a tiny bit forward, or maybe it's the car turning its wheels. So, every "beep" is like one step or action the toy car does.
Clock Frequency: Now, let's say the car is very fast and it can do 3 "beeps" every second. That means it's doing its actions very quickly! So, the clock frequency is like how fast the toy car can do its actions in one second. If it can do 3 "beeps" per second, we say its clock frequency is 3 times per second.

Clock cycles are indeed like missions or tasks that the processor needs to complete, such as fetching instructions, decoding them, and executing them. Clock frequency, on the other hand, is like how often the processor can perform those missions or tasks in a given amount of time, usually measured in cycles per second (hertz). So, higher clock frequency means the processor can complete those missions more frequently, which often translates to faster performance.

==============================================================================================================================================================================================================

**The Von Neumann Bottleneck**
Imagine you're working on a project in your room, and you have a bookshelf across the hall where you store all your materials. Let's say you need to frequently reference information from this bookshelf while working on your project.

Here's how the Von Neumann bottleneck relates to this scenario:

- Your Work (Processing Unit): You represent the processing unit, which is actively working on your project—reading, writing, and solving problems.

- Bookshelf (Memory Unit): The bookshelf across the hall represents the memory unit, where you store all your materials and information related to your project.

- Walking to the Bookshelf (Data Fetching): Every time you need to reference information from the bookshelf, you have to walk across the hall to fetch it. This represents the time it takes for the processor to access data from memory.

- Walking Back to Work (Data Transfer): After fetching the information, you walk back to your workspace to continue working on your project. This back-and-forth movement represents the transfer of data between memory and the processing unit.

- Time Spent Walking (Bottleneck): If you find yourself spending more time walking to and from the bookshelf than actually working on your project, it becomes a bottleneck. 

Similarly, in computer systems, if the processor spends more time waiting for data from memory than executing instructions, it slows down overall performance, creating the Von Neumann bottleneck.

==============================================================================================================================================================================================================

**Moore's Law** is the observation made by Gordon Moore, co-founder of Intel, in 1965, that the number of transistors on a microchip doubles approximately every two years, leading to a continuous increase in computing power and performance, and a decrease in cost. Essentially, it predicts that the capabilities of computer technology will exponentially grow over time, resulting in faster and more powerful devices.
