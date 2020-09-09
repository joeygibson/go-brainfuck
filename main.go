package main

import (
	"fmt"
	"go-brainfuck/brainfuck"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]
	code, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}

	compiler := brainfuck.NewCompiler(string(code))
	instructions := compiler.Compile()

	machine := brainfuck.NewMachine(instructions, os.Stdin, os.Stdout)
	machine.Execute()
}
