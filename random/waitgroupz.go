package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printhello() {
	wg.Add(1)
	defer wg.Done()
	fmt.Println("hello world")
}

func main() {

	for i := 1; i <= 5; i++ {
		go printhello()
	}

	wg.Wait()

	fmt.Println("all goroutines finished")
}
