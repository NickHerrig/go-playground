// echo prints the command line arguments
package main

import (
	"fmt"
	"os"
	"strings"
)

// exercise 1.1: Modify echo to print os.Args[0]
func one() string {
	return strings.Join(os.Args, " ")
}

// exercise 1.2 Modify ech to print index/value , on per line
func two() string {

	var line string
	for i, arg := range os.Args {
		line += fmt.Sprintf("%d %s\n", i, arg)
	}
	return line

}

func main() {
	fmt.Println(one())
	fmt.Println(two())

}
