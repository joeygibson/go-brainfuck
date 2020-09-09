package main

import (
	"fmt"
	"go-brainfuck/machine"
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

	m := machine.NewMachine(string(code), os.Stdin, os.Stdout)
	m.Execute()
}
