package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    start := time.Now()
    for i, arg := range os.Args[:] {
        fmt.Println("Arg:", i, "value:", arg)
    }
    fmt.Println(time.Since(start))
}
