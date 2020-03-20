package main

import "fmt"

func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

func main() {
    fmt.Println(add(1,2,3,4,5))
}

// Similar to the *args and **kwargs in python
// def add (*args):
//     return sum(args)
// 
// print(add(1,2,3,4))

