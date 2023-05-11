package main

import (
	"flag"
	"fmt"
	"os"
	"otus-architecture/hw06/program"
)

var (
	input  string
	output string
)

func init() {
	flag.StringVar(&input, "input", "", "Path to input file")
	flag.StringVar(&output, "output", "", "Path to output file")
}

func main() {
	flag.Parse()

	p2 := program.Program2{}

	if err := p2.Write2MatrixToFile(input); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := p2.ExecuteMatrixAdd(input, output); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
