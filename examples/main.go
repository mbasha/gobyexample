package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the example name to run (e.g., 'hello-world').")
		return
	}

	example := os.Args[1]

	switch example {
	case "hello-world":
		HelloWorld()
	case "if-else":
		IfElse()
	case "switch":
		Switch()
	case "for":
		For()
	case "values":
		Values()
	case "variables":
		Variables()
	case "constant":
		Constant()
	case "arrays":
		Arrays()
	case "slices":
		Slices()
	case "maps":
		Maps()
	case "functions":
		Functions()
	case "multiple-return-values":
		MultipleReturnValues()
	case "variadic-functions":
		VariadicFunctions()
	default:
		fmt.Printf("Unknown example: %s\n", example)
	}
}
