package main

import (
    "fmt"
    "os"
    "time"
    "strings"
)

func runInefficient() {
    start := time.Now()
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println("INEFFICIENT")
    fmt.Println(s)
    fmt.Println(time.Since(start))
}

func runEfficient() {
    start := time.Now()
    fmt.Println("EFFICIENT")
    fmt.Println(strings.Join(os.Args[1:], " "))
    fmt.Println(time.Since(start))
}
func main() {
    runInefficient()
    runEfficient()
}
