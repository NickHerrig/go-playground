package main

// Idiomatic go has import statements like this.
import (
    "fmt"
    "reflect"
)

func main() {
    x := 1
    fmt.Println(reflect.TypeOf(x))
    fmt.Println("Hello, World!")
}

