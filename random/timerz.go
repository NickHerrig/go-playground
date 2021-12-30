package main

import (
	"fmt"
	"time"
)


func coro(interval time.Duration) {
	c := time.Tick(interval)
	for {
		fmt.Println("Hello from coro")
		<-c
	}
}


func main() {
	go coro(5 * time.Second)
	time.Sleep(5 * time.Second)
	go coro(1 * time.Second)
	select{}
}
