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

	mainArgument := strings.ToLower(argWithProg[1])

	switch mainArgument {
	case "add":
		fmt.Println("add received")
	case "list":
		fmt.Println("list received")
	case "summary":
		//Can have --month, --id argument or none
		fmt.Println("list received")
	case "delete":
		fmt.Println("list received")
	case "":
		fmt.Println("no needed")

	}
}
