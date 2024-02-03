package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file := "input.txt"
	/*if len(os.Args) == 1 {
		file = "sample" + file
	} else { // go run . [filepath]
		file = os.Args[1] + file
	}*/
	fmt.Println(file)
	input := ReadInput(file)
	Part1(input)
	Part2(input)
}
func ReadInput(filepath string) (input []string) {
	file_in, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file_in.Close()
	scanner := bufio.NewScanner(file_in)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
