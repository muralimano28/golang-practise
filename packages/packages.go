package main

import (
    "fmt"
    "math/rand"
    "github.com/muralimano28/golang-practise/hello"
)

func main() {
    // Every go program is made up of packages
    // Program starts running in package main
    // Here am importing two packages. 1. fmt - format 2. math/rand - for generating random number
    // And files inside math/rand will have package name as 'rand'.
    fmt.Println("This line is from math/rand package, my fav random number: ", rand.Intn(10))
    fmt.Println("Below line is printed from hello package that I created")
    // fmt.Println(hello.Dummy1())
    // Above line wont work because, hello.Dummy1() doesnt return anything.
    // package.<function name> this function name should start in capital letter, else it wont work
    hello.Dummy1()
    hello.Dummy2()
}
