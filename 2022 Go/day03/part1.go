package main

import (
	"fmt"
	"strings"
)

func Part1(input []string) {
	rksk := Parse1(input)
	sum := 0
	for _, r := range rksk {
		seen := make(map[byte]struct{})
		for i := range r[0] {
			if _, f := seen[r[0][i]]; !f && strings.Contains(r[1], string(r[0][i])) {
				seen[r[0][i]] = struct{}{}
				sum += Rtoi(r[0][i])
			}
		}
	}
	fmt.Println(sum)
}

func Parse1(input []string) (rucksack [][2]string) {
	for _, line := range input {
		rucksack = append(rucksack, [2]string{line[:len(line)/2], line[len(line)/2:]})
	}
	return
}

func Rtoi(r byte) int {
	if r >= 'a' && r <= 'z' {
		return int(r-'a') + 1
	} else if r >= 'A' && r <= 'Z' {
		return int(r-'A') + 27
	}
	return 0
}
