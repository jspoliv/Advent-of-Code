package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type coord struct {
	x, y int
}

func main() {
	file_in, _ := os.Open("input.txt")
	defer file_in.Close()
	input := parseInput(file_in)
	//part1(input)
	part2(input)
}

func parseInput(file_in *os.File) []string {
	scanner := bufio.NewScanner(file_in)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func part2(input []string) {
	emptyX, emptyY := part2Parse(input)
	var node []coord
	for i, line := range input {
		for j, char := range line {
			if char == '#' {
				node = append(node, coord{i, j})
			}
		}
	}
	sum := 0
	for i := 0; i < len(node); i++ {
		for j := 0; j < len(node); j++ {
			if j > i {
				sum += mhtDist(node[i], node[j])
				for _, x := range emptyX {
					if (node[i].x < x && x < node[j].x) || (node[i].x > x && x > node[j].x) {
						sum += 1000000 - 1
					}
				}
				for _, y := range emptyY {
					if (node[i].y < y && y < node[j].y) || (node[i].y > y && y > node[j].y) {
						sum += 1000000 - 1
					}

				}
			}
		}
	}
	fmt.Println(sum)
}

func part2Parse(input []string) (emptyX, emptyY []int) {
	emptyX = checkEmpty(input)
	tmp := transpStr(input)
	emptyY = checkEmpty(tmp)
	return
}

func checkEmpty(input []string) []int {
	var empty []int
	for i, line := range input {
		if !strings.ContainsRune(line, '#') {
			empty = append(empty, i)
		}
	}
	return empty
}

func part1(input []string) {
	tmp := part1Parse(input)
	var node []coord
	for i, line := range tmp {
		for j, char := range line {
			if char == '#' {
				node = append(node, coord{i, j})
			}
		}
	}
	sum := 0
	for i := 0; i < len(node); i++ {
		for j := 0; j < len(node); j++ {
			if j > i {
				sum += mhtDist(node[i], node[j])
			}
		}
	}
	fmt.Println(sum)
}

func mhtDist(a, b coord) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func part1Parse(input []string) []string {
	tmp := doubleEmpty(input)
	tmp = transpStr(tmp)
	tmp = doubleEmpty(tmp)
	tmp = transpStr(tmp)
	return tmp
}

func doubleEmpty(input []string) []string {
	var empty []int
	var empty_line string
	for i, line := range input {
		if !strings.ContainsRune(line, '#') {
			empty = append(empty, i)
			empty_line = line
		}
	}
	for i := 0; i < len(empty); i++ {
		t := empty[i]
		if i+1 < len(empty) {
			for j := i; j < len(empty); j++ {
				empty[j]++
			}
		}
		left := input[:t]
		right := input[t:]
		input = append(left, empty_line)
		input = append(input, right...)
	}
	return input
}

func transpStr(input []string) []string {
	transp := make([][]rune, len(input[0]))
	for _, line := range input {
		for j, char := range line {
			transp[j] = append(transp[j], char)
		}
	}
	var tmp []string
	for _, line := range transp {
		tmp = append(tmp, string(line))
	}
	return tmp
}
