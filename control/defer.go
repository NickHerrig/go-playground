package main

import "fmt"

func first(){
    fmt.Println("First")
}


func second(){
    fmt.Println("SECOND")
}


func main(){
    defer second()
    first()
}

// defer is usually used for closing files. 
// defer happens even if a panci happens
