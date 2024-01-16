package main

import (
	"fmt"
)

const (
	ROCK int = iota + 1
	PAPER
	SCISSORS
	A        = ROCK
	B        = PAPER
	C        = SCISSORS
	X        = ROCK
	Y        = PAPER
	Z        = SCISSORS
	LOSS int = 0
	DRAW int = 3
	WIN  int = 6
)

func Part1(input []string) {
	pairs := Parse1(input)
	sum := 0
	for _, p := range pairs {
		sum += RPS1(p[0], p[1])
	}
	fmt.Println(sum)
}

func Parse1(input []string) [][2]int {
	var pair [][2]int
	for _, line := range input {
		var p [2]int
		switch line[0] {
		case 'A':
			p[0] = ROCK
		case 'B':
			p[0] = PAPER
		case 'C':
			p[0] = SCISSORS
		}
		switch line[2] {
		case 'X':
			p[1] = ROCK
		case 'Y':
			p[1] = PAPER
		case 'Z':
			p[1] = SCISSORS
		}
		pair = append(pair, p)
	}
	return pair
}

func RPS1(a, b int) int {
	if a == b {
		return b + DRAW
	}
	if (b == ROCK && a == SCISSORS) || (b == PAPER && a == ROCK) || (b == SCISSORS && a == PAPER) {
		return b + WIN
	}
	return b + LOSS
}
