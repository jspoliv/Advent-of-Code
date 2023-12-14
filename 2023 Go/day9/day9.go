package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file_in, _ := os.Open("input.txt")
	defer file_in.Close()
	input := parseInput(file_in)
	//fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func parseInput(file_in *os.File) []string {
	scanner := bufio.NewScanner(file_in)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func part1Parse(input []string) [][]int {
	var fields [][]int
	for _, line := range input {
		str_fields := strings.Fields(line)
		var int_fields []int
		for _, field := range str_fields {
			t, _ := strconv.Atoi(field)
			int_fields = append(int_fields, t)
		}
		fields = append(fields, int_fields)
	}
	return fields
}

func part1(input []string) (sum int) {
	fields := part1Parse(input)
	sum = 0
	for _, line := range fields {
		sum += nextVal(line)
	}
	return
}

func part2(input []string) (sum int) {
	fields := part1Parse(input)
	sum = 0
	for _, line := range fields {
		sum += prevVal(line)
	}
	return
}

func nextVal(history []int) int {
	var stability [][]int
	stability = append(stability, history)
	for loop := 0; loop < len(history); loop++ {
		var tmp []int
		for i := 0; i < len(stability[loop])-1; i++ {
			tmp = append(tmp, stability[loop][i+1]-stability[loop][i])
		}
		if allZeros(tmp) {
			break
		} else {
			stability = append(stability, tmp)
		}
	}
	result := 0
	for i := len(stability) - 1; i >= 0; i-- {
		result = result + stability[i][len(stability[i])-1]
	}
	return result
}

func prevVal(history []int) int {
	var stability [][]int
	stability = append(stability, history)
	for loop := 0; loop < len(history); loop++ {
		var tmp []int
		for i := 0; i < len(stability[loop])-1; i++ {
			tmp = append(tmp, stability[loop][i+1]-stability[loop][i])
		}
		if allZeros(tmp) {
			break
		} else {
			stability = append(stability, tmp)
		}
	}
	result := 0
	for i := len(stability) - 1; i >= 0; i-- {
		result = stability[i][0] - result
	}
	return result
}

func allZeros(integers []int) bool {
	result := true
	for i := 0; i < len(integers); i++ {
		if integers[i] != 0 {
			result = false
			break
		}
	}
	return result
}
