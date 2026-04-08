package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		fmt.Println("Usage: concourse <config-file>")
		os.Exit(1)
	}

	// Print what we see
	fmt.Printf("Concourse received config file: %v\n", argsWithoutProg)
}
