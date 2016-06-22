package main

import (
    "fmt"
)

// Below add function takes two parameters of type int.
// Note that type comes after the variable name.
// If both the variables share same type, then we can omit types for all except the last one.
// For example check below sub function.
func sub (x, y int) int {
    return x - y
}

// type at last of the funtion is the return type of this function
func add (x int, y int) int {
    return x + y
}

// A function can return multiple values
// For example below swap function returns two string.
func swap (x, y string) (string, string) {
   return y, x 
}

// A function can return named variables. If so, they are named at the top of the function.
func split (sum float32) (x, y float32) {
    x = sum / 2
    y = x
    // Here x and y are returned.
    // This is called naked return and it should be used only in short functions. Because, it will hinder readability in long functions.
    return 
}

func main() {
    // Export functions should start in capital letters and non export functions should start in small letters.
    fmt.Println("This file contains my first function that adds two numbers")
    fmt.Println("Sum of 2 and 4 is: ", add(2, 4))
    fmt.Println("Difference of 4 and 2 is ", sub(4, 2))
    a, b := swap("murali", "manohar")
    fmt.Println("Swapped result of 'murali' and 'manohar' is ", a, b)
    c, d := split(17)
    fmt.Println("Splited values of 17 is ", c, d) 
    // Lesson learnt:
    // 1. You can run only main packages with main function. And main function is the entry point to run go code.
    // 2. Go code with other package names is not compiled while executing "go install"
    //    Instead they are imported by other go code and compiled with that code and executable is created for that code
}
