package main

import (
	"os"
	"strings"
)

func main() {
	EchoRange()
	EchoJoin()
}

func EchoRange() {

	for index, value := range os.Args {
		s := []interface{}{index, value}
		s = append(s, s)
	}

}

func EchoJoin() {
	_ = strings.Join(os.Args, " ")

}
