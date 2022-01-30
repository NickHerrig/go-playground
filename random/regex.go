package main

import (
	"fmt"
	"regexp"
)

var (
	ZipcodeRX = regexp.MustCompile("^\\d{1,5}$")
)

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func main() {
	fmt.Println(Matches("hello", ZipcodeRX))
	fmt.Println(Matches("world", ZipcodeRX))
	fmt.Println(Matches("500", ZipcodeRX))
	fmt.Println(Matches("9", ZipcodeRX))
	fmt.Println(Matches("900", ZipcodeRX))
	fmt.Println(Matches("50014", ZipcodeRX))
	fmt.Println(Matches("500141", ZipcodeRX))
}
