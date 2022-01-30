package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	b := make([]byte, 16)
	n, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(n)
	fmt.Println(b)
}
