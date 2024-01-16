package main

import "fmt"

func Parse2(input []string) [][2]int {
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
			p[1] = LOSS
		case 'Y':
			p[1] = DRAW
		case 'Z':
			p[1] = WIN
		}
		pair = append(pair, p)
	}
	return pair
}

func Part2(input []string) {
	pairs := Parse2(input)
	sum := 0
	for _, p := range pairs {
		sum += RPS2(p[0], p[1])
	}
	fmt.Println(sum)
}

func RPS2(a, b int) int {
	if b == LOSS {
		if a == ROCK {
			return LOSS + SCISSORS
		} else if a == PAPER {
			return LOSS + ROCK
		} else { //if a == SCISSORS
			return LOSS + PAPER
		}
	} else if b == DRAW {
		if a == ROCK {
			return DRAW + ROCK
		} else if a == PAPER {
			return DRAW + PAPER
		} else { //if a == SCISSORS
			return DRAW + SCISSORS
		}
	} else { //if b == WIN
		if a == ROCK {
			return WIN + PAPER
		} else if a == PAPER {
			return WIN + SCISSORS
		} else { //if a == SCISSORS
			return WIN + ROCK
		}
	}
}
