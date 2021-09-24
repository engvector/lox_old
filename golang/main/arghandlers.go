package main

import "fmt"

// Start an interactive interpreter
func start_interpreter() {
	fmt.Println(Usage())
	fmt.Println("Starting interpreter...")
	fmt.Println("> ")
}

// Validate subcommand and execute
func handle_subcommand(subcommand string, args []string) {
	switch subcommand {
		case "run":
			fmt.Println("Running lox file -", args[0])
		case "build":
			fmt.Println("Building...")
		default:
			fmt.Println("ERROR!!! Bad subcommand:", subcommand)
	}
}

// Run the program
func run(args []string) {
}

// Build an application/module
func build(args []string) {
}
