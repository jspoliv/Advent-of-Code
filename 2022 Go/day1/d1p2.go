package main

import (
	"fmt"
)

func Part2(input []string) {
	elf := Parse1(input)
	hi1, hi2, hi3 := -1, -1, -1
	for _, inv := range elf {
		sum := 0
		for _, food := range inv {
			sum += food
		}
		if sum > hi1 {
			hi3 = hi2
			hi2 = hi1
			hi1 = sum
		} else if sum > hi2 {
			hi3 = hi2
			hi2 = sum
		} else if sum > hi3 {
			hi3 = sum
		}
	}
	fmt.Println(hi1 + hi2 + hi3)
}
