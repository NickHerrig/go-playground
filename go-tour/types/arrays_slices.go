package main

import (
    "fmt"
)

func main() {
    my_array := [6]int{45,645,7,345,56,435}
    for _, value := range my_array {
        fmt.Println(value)
    }
    // initiated an empty slice
    var my_slice []int
    fmt.Println(my_slice)

    // inite an slice from an array
    my_slice_init := my_array[0:5]
    fmt.Println(my_slice_init)

    // a slice of length 5 with a capacity of 10
    my_slice_capacity := make([]int, 5, 10)
    fmt.Println(my_slice_capacity)

    // append 4 and 6 to my slice
    appended := append(my_slice_init, 4, 6)
    fmt.Println(appended)

    // copy slice src to slice dst.
    coppied := copy(my_slice_capacity, appended)
    fmt.Println(coppied)
}
