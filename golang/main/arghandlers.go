package main

import (
				"fmt"
				"os")

// Start an interactive interpreter
func Interpreter() {
	fmt.Println(Usage())
	fmt.Println("Starting interpreter...")
	fmt.Println("> ")
}

// Run the program
func Run(args []string) {
}

// Build an application/module
func Build(args []string) {
}

// Error reporting functions
func ReportError(line int, message string) {
	Report(line, "", message)
}

func Report(line int, where string, message string) {
	fmt.Fprintln(os.Stderr,
							"ERROR @",line,":",where,"-",message);
}
