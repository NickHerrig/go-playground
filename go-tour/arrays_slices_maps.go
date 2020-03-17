package main

import (
    "fmt"
)

func main() {
    var my_array [6]int
    my_array[3] = 100
    for _, value := range my_array {
        fmt.Println(value)
    }
}
