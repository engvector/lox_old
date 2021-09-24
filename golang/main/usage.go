package main

var usageString string = 
`
Lox interpreter / compiler
--------------------------
NOTE: Language syntax is explained in Crafting Interpreters book (https://craftinginterpreters.com/). This implementation differs from the Java version present in that book.

Interactive use:
	$ lox
	> 
Non-Interactive use:
	$ lox run <path>
	$ lox build .
	$ lox build -o <binary> <path>
	$ lox build -m <module> <path>
`

func Usage() string {
	return usageString
}
