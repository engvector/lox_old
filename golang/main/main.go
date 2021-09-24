package main

import ("fmt"
				"os"
				/*"lox/scanner"*/)

func main() {
	args := os.Args[1:]

	switch {
		case len(args) == 0:		start_interpreter()
		case len(args) > 1:			handle_subcommand(args[0], args[1:])
		default:								fmt.Println(Usage())
	}
}

