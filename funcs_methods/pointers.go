package main

import "fmt"

func add(my_pointer *int) {
    *my_pointer = 0
}

func main() {
    x := 11
    add(&x)
    fmt.Println(x)

}
