package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {

    for _, url := range os.Args[1:] {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprint(os.Stderr, err)
            os.Exit(1)
        }
        _, err = io.Copy(os.Stdout, resp.Body)
        if err != nil {
            fmt.Fprint(os.Stderr, err)
            os.Exit(1)
        }
    }
}
