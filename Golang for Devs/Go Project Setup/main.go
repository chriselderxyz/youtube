package main

import (
	"fmt"

	"github.com/chriselderxyz/my-project/printer"
	"github.com/fatih/color"
)

func main() {
	// Entry point for our Go application
	fmt.Println("Hello, World!")
	printer.PrintPublic("Hello, World!")
	color.Cyan("Hello, World!")
}
