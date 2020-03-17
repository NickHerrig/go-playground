package main

import (
    "fmt"
)

func main() {
// All of this code shows experiments with arrays and slices. 
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

// All of this code shows experiments with Maps (Go's dictionary, hashtable, associate array)
      // Example of a runtime error!
//    var my_map map[string]int
//    my_map["key"] = 11
//    fmt.Println(my_map)

    elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["O"] = "Oxygen"
    fmt.Println(elements)
    for key, value := range elements {
        fmt.Println("The element symbol:",key,"stands for", value)
    }

}
