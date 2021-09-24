package main

import ("fmt"
				"os"
				/*"lox/scanner"*/)

func main() {
	args := os.Args[1:]

	switch {
		case len(args) == 0:		Interpreter()
		case len(args) > 1:			handle_subcommand(args[0], args[1:])
		default:								fmt.Println(Usage())
	}
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

