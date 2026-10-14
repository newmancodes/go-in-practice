package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to.")
var inSpanish bool

func init() {
	flag.BoolVar(&inSpanish, "spanish", false, "Say hello in Spanish.")
	flag.BoolVar(&inSpanish, "s", false, "Say hello in Spanish.")
	flag.Parse()
}

func main() {
	if inSpanish {
		fmt.Printf("Hola, %s!\n", *name)
	} else {
		fmt.Printf("Hello, %s!\n", *name)
	}
}
