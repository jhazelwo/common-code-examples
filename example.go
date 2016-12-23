package main

// Import a module
import "fmt"

// Define a variable
var hi = "hello world"


// Print "hello world"
fmt.Println("hello world")


// If, then, else
if 1 == 1 {
    fmt.Println(num, "is negative")
} else if 2 == 2 {
    fmt.Println(num, "has 1 digit")
} else {
    fmt.Println(num, "has multiple digits")
}

// Iterate
for i := 1;  i<=2; i++ {
    s += i
}


// Define a function
func foo(i int) {
    return i
}

// Run a function
foo(2)

// Define a class
type car struct {
    make string
    year int
}

// Instantiate a class
car{make: "abc", year: 2000}
