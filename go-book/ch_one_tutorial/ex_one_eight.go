package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {

    for _, url := range os.Args[1:] {
        if ! strings.HasPrefix(url, "https://") {
            url = "https://" + url
        }

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
        fmt.Println("The Response Code is:", resp.Status)
    }
}
