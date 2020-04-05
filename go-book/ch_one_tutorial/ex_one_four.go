package main

import (
    "bufio"
    "fmt"
    "os"
)

func countLines(f *os.File, counts map[string]map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        if counts[input.Text()] == nil {
            counts[input.Text()] = make(map[string]int)
        }
        counts[input.Text()][f.Name()]++
    }
}

func main() {
    counts := make(map[string]map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else{
        for _, arg := range files {
            f, err :=  os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "ERROR: %v/n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }
    for word, files := range counts {
        fmt.Println(word)
        fmt.Println(files)
    } 
}
