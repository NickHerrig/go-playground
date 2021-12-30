package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

func main() {

	type config struct {
		Key1 string
		Key2 string
		Key4 bool
	}

	var conf config
	md, err := toml.DecodeFile("file.toml", &conf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Undecoded keys: %q\n", md.Undecoded())
}
