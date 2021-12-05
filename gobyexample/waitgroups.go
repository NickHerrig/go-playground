package main

import (
	"fmt"
	"sync"
	"time"
)

func hello(name string, sleep time.Duration) {
	fmt.Printf("Hello %s!\n", name)
	time.Sleep(sleep)
	fmt.Printf("Hello again %s!\n", name)
}

func main() {
	var wg sync.WaitGroup
	friends := map[string]time.Duration{
		"nick":  1 * time.Second,
		"megan": 2 * time.Second,
		"jeff":  5 * time.Second,
	}

	for name, time := range friends {
		wg.Add(1)

		// https://go.dev/doc/faq#closures_and_goroutines
		name, time := name, time
		go func() {
			defer wg.Done()
			hello(name, time)
		}()
	}

	wg.Wait()

}
