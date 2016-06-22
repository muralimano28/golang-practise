package main

import (
    "fmt"
)

// "var" statement declares a variable. As in function arguments, the data type is specified at last.
// "var" statement can be declared at function level and package level. In below example, variable i is declared at package level and variable j is declared at function level.

var i int

func main () {
    var j int
    // We can declared multiple variables at a time and if all these variables have same data types then we can mention the data type at last.
    var k, l, m bool
    // Go lang doesnt allow to declare a variable that is not used anywhere in the function.
    // But package level variable declaration allows to declare variable that is not used.
    // For example remove i from below line and it wont throw an error.
    fmt.Println("Value of i and j is ", i, j)
    fmt.Println("Value of k, l and m is ", k, l, m)

    // Variable with initializers
    // 1. var declaration can include intializers one per variable
    // 2. If initializer is present, then we can omit the type. variable will take the type of initializer.
    var a, b int = 1, 2
    var c, d = 3, 4
    fmt.Println("Value of a and b is ", a, b)
    fmt.Println("Value of c and d is ", c, d)

    // Short variable declarations
    // 1. Inside function, the := short assignment can be used in place of var declaration with implicit type.
    // Check below example.
    e, f := 5, "hello"
    fmt.Println("Value of e and f is ", e, f)
    // 2. Outside function every line should start with either func or var and so on. So, short variable declaration is not allowed.
}
