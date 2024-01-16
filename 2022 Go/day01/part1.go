package main

import (
	"fmt"
	"strconv"
)

func Part1(input []string) {
	elf := Parse1(input)
	hi := -1
	for _, inv := range elf {
		sum := 0
		for _, food := range inv {
			sum += food
		}
		if sum > hi {
			hi = sum
		}
	}
	fmt.Println(hi)
}

func Parse1(input []string) [][]int {
	var elf [][]int
	var tmp []int
	for _, line := range input {
		if line != "" {
			t, _ := strconv.Atoi(line)
			tmp = append(tmp, t)
		} else {
			elf = append(elf, tmp)
			tmp = nil
		}
	}
	elf = append(elf, tmp)
	return elf
}
