package main

import(
    "fmt"
    "os"
    "io/ioutil"
)

func readdirectory() {
    dir, err := os.Open(".")
    if err != nil {
        panic(err)
    }
    defer dir.Close()

    fileInfo, err := dir.Readdir(-1)
    if err != nil {
        panic(err)
    }
    for _, fi := range fileInfo {
        fmt.Println(fi.Name())
    }
}

func createfile() {
    file, err := os.Create("test.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    file.WriteString("Hello, World\n")
}

func longread() {
    file, err := os.Open("test.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    stat, err := file.Stat()
    if err != nil {
        panic(err)
    }

    bytes := make([]byte, stat.Size())
    _, err = file.Read(bytes)
    if err != nil {
        panic(err)
    }

    str := string(bytes)
    fmt.Println("long code:", str)
}


func shortread() {
    bytes, err := ioutil.ReadFile("test.txt")
    if err != nil {
        panic(err)
    }

    str := string(bytes)
    fmt.Println("short code:", str)
}

func main() {
    readdirectory()
}
