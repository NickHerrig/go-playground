package main

import (
    "fmt"
    "strings"
)

// The strings packages https://golang.org/pkg/strings/#pkg-overview

func main() {

    fmt.Println(strings.Join([]string{"Hello", "World"}, ", "))

    fmt.Println(strings.Split("Hello, World", ", "))

    fmt.Println(strings.ToUpper("Hello, World"))

 
}
