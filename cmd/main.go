package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/placeholder30/wcl/pkg/file"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("we need args bro pass -h for help")
		os.Exit(1)
	}
	if strings.ToLower(os.Args[1]) == "-v" {
		fmt.Println("v 0.0.1")
		os.Exit(0)

	}
	if strings.ToLower(os.Args[1]) == "-h" {
		fmt.Println("welcome!")
		fmt.Println("pass me a file or files as args and i'll tell you some stuff about it.")
		os.Exit(1)
	}
	file.File(os.Args[1:])
	os.Exit(0)
}
