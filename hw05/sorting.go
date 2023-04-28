package main

import (
	"flag"
	"fmt"
	"os"
	"otus-architecture/hw05/ioc"
	"strconv"
	"strings"
)

var (
	sortMethod string
	inputFile  string
	outputFile string
)

func init() {
	flag.StringVar(&sortMethod, "method", "", "Sort method")
	flag.StringVar(&inputFile, "input", "", "Path to input file")
	flag.StringVar(&outputFile, "output", "", "Path to output file")
}

func main() {
	flag.Parse()

	data, err := readArrayFromFile(inputFile)
	if err != nil {
		fmt.Println("Error while reading file:", err)
		os.Exit(1)
	}

	factory := ioc.New()
	command := factory.Resolve(sortMethod, data)

	if command == nil {
		fmt.Println("Unknown Sort Method:", sortMethod)
		os.Exit(1)
	}

	err = command.Execute()
	if err != nil {
		fmt.Println("Error while sorting:", err)
		os.Exit(1)
	}

	err = writeOutputFile(outputFile, sortMethod, data)
	if err != nil {
		fmt.Println("Error while writing file:", err)
		os.Exit(1)
	}
}

func readArrayFromFile(inputFile string) ([]int, error) {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	numbers := strings.Split(string(content), "\n")

	var arr []int

	for idx := range numbers {
		if numbers[idx] == "" {
			continue
		}

		n, err := strconv.Atoi(numbers[idx])
		if err != nil {
			return nil, err
		}
		arr = append(arr, n)
	}

	return arr, nil
}

func writeOutputFile(outputFile string, method string, arr []int) error {
	result := strings.Builder{}
	result.WriteString("Sort Method: " + method + "\n")

	for idx := range arr {
		result.WriteString(strconv.Itoa(arr[idx]) + "\n")
	}

	return os.WriteFile(outputFile, []byte(result.String()), 0666)
}
