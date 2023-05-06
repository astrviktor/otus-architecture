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

	p1 := program.Program1{}
	err := p1.ExecuteMatrixAdd(input, output)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
