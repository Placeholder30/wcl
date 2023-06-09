package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/placeholder30/wcl/pkg/utils"
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
		fmt.Println("pass me a file or a list of files and i'll tell you stuff about `em.")
		os.Exit(0)
	}
	utils.File(os.Args[1:])
	os.Exit(0)
}
