package main

import "fmt"

func main() {

// the traditional for loop
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)

// the while statement is simply a for loop...
    sum = 1
    for sum < 1000 {
        sum += sum 
    }
    fmt.Println(sum)
}
