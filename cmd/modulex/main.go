package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"log"
)

// run `go install .` inside this file's directory
// then run: modulex
// app should be added in %GOPATH%/bin, see: `ll $GOPATH/bin/modulex`
func main() {
	fmt.Println("Enter Your Name: ")

	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello Mr. %s!\n", name)

	fmt.Println(cmp.Diff(name, "12345"))
}
