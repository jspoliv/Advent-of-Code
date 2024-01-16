package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file_in, _ := os.Open("input.txt")
	defer file_in.Close()
	node_map, inst := parseInput(file_in)
	fmt.Println(part2(node_map, inst))
}

func part2(node_map map[string][2]string, instr []int) int {
	var endA []string
	for key := range node_map {
		if strings.LastIndex(key, "A") == len(key)-1 {
			endA = append(endA, key)
		}
	}
	var cycle []int
	for _, e := range endA {
		cycle = append(cycle, part1(e, node_map, instr))
	}

	return LCM(cycle[0], cycle[1], cycle[1:]...)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

func part1(current string, node_map map[string][2]string, inst []int) (steps int) {
	inst_len := len(inst)
	steps = 0
	for ; current[len(current)-1:] != "Z"; steps++ { // current != "ZZZ" <-- part1
		current = node_map[current][inst[steps%inst_len]]
	}
	return
}

func parseInput(file_in *os.File) (map[string][2]string, []int) {
	scanner := bufio.NewScanner(file_in)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	node_map := make(map[string][2]string)
	for i := 2; i < len(input); i++ {
		tmp, _ := strings.CutSuffix(input[i], ")")
		label, tmp, _ := strings.Cut(tmp, " = (")
		neighbors := strings.Split(tmp, ", ")
		node_map[label] = [2]string{neighbors[0], neighbors[1]}
	}

	return node_map, parseInstr(input[0])
}

func parseInstr(in string) (out []int) {
	for _, x := range in {
		if x == 'L' {
			out = append(out, 0)
		} else {
			out = append(out, 1)
		}
	}
	return
}
