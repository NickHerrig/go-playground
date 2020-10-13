package main

import (
    "fmt"

    "github.com/nickherrig/go-playground/greetings"
)

func main() {
    message := greetings.Hello("Nick")
    fmt.Println(message)
}
