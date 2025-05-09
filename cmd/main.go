package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	argWithProg := os.Args

	if len(argWithProg) < 2 {
		fmt.Println("Plesa add the right arguments ")
		return
	}

	argument := strings.ToLower(argWithProg[1])
	switch argument {
	case "add":
		fmt.Println("add received")
	case "list":
		fmt.Println("list received")

	case "":
		fmt.Println("no needed")

	}
}
