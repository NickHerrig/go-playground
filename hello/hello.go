package main

import (
    "fmt"
    "log"

    "github.com/nickherrig/go-playground/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    message, err := greetings.Hello("Nick")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}
