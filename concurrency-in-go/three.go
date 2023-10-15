package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func sameAddrSpace(){
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
}

func tricky(){
	var wg sync.WaitGroup
	for _, salutation :=range []string{"hello", "greetings", "goodbye"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

func fixed(){
	var wg sync.WaitGroup
	for _, salutation :=range []string{"hello", "greetings", "goodbye"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}

func memAlloc(){
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}
	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c }
	
	const numGoroutines =1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i>0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000 )

}

func runWg() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func () {
		defer wg.Done()
		fmt.Println("First goroutine sleeping")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func () {
		defer wg.Done()
		fmt.Println("second goroutine sleeping")
		time.Sleep(2)
	}()
	wg.Wait()
	fmt.Print("All goroutines finished")

}




func main() {
	//sameAddrSpace()
	//tricky()
	//fixed()
	//memAlloc()
	//runWg()


	hello := func (wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v\n", id)
	}

	var wg sync.WaitGroup
	numRuns := 5
	wg.Add(numRuns)
	for i := 0; i < numRuns; i++ {
		hello(&wg, i+1)
	}
	wg.Wait()

}