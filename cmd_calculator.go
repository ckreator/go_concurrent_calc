package main

import (
	"cmd_calculator/interpreter"
	"fmt"
)

func main() {
	//s := "5 + 14.3e-1 * (14 + (16 - -(-(9))))"
	s := "1 - 4 + 9"
	// RPN
	// 5 14.3e-1 + 14 16 - neg neg 9 + *
	fmt.Println("Running in main => ", s)
	interpreter.Interpreter(s)
}
