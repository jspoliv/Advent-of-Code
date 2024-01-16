package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := ".txt"
	if len(os.Args) == 1 {
		file = "sample" + file
	} else { // go run . [filepath]
		file = os.Args[1] + file
	}
	fmt.Println(file)
	input := ParseInput(file)
	Part1(input)
	Part2(input)
}

func ParseInput(filepath string) (input []string) {
	file_in, _ := os.Open(filepath)
	defer file_in.Close()
	scanner := bufio.NewScanner(file_in)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
