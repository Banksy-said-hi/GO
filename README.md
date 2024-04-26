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

```var i interface{} = 42
var j interface{} = "hello"
var k interface{} = 3.14```
