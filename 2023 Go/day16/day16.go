package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos_t struct {
	i, j int
	dir  string
}

var WALL pos_t = pos_t{-1, -1, "X"}

func main() {
	file := ".txt"
	if len(os.Args) == 1 {
		file = "sample" + file
	} else {
		file = os.Args[1] + file
	}
	file_in, _ := os.Open(file)
	fmt.Println(file)
	defer file_in.Close()
	input := ParseInput(file_in)
	fmt.Println(Part1(input, pos_t{0, 0, "E"}))
	Part2(input)
}

func ParseInput(file_in *os.File) (input [][]rune) {
	scanner := bufio.NewScanner(file_in)
	for scanner.Scan() {
		input = append(input, []rune(scanner.Text()))
	}
	return input
}

func Part1(input [][]rune, e pos_t) int {
	m := Part1Parse(input, e)
	walked := make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		walked[i] = make([]int, len(input[i]))
	}

	for k := range m {
		walked[k.i][k.j]++
	}
	sum := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if walked[i][j] > 0 {
				sum++
			}
		}
	}
	return sum
}

func Part2(input [][]rune) {
	high := 0

	for i := 0; i < len(input); i++ {
		tmp := Part1(input, pos_t{i, 0, "E"})
		if tmp > high {
			high = tmp
		}
		tmp = Part1(input, pos_t{i, len(input[i]) - 1, "W"})
		if tmp > high {
			high = tmp
		}
	}

	for j := 0; j < len(input[0]); j++ {
		tmp := Part1(input, pos_t{0, j, "S"})
		if tmp > high {
			high = tmp
		}
		tmp = Part1(input, pos_t{len(input) - 1, j, "N"})
		if tmp > high {
			high = tmp
		}
	}

	fmt.Println(high)
}

func findPos(input [][]rune, e pos_t) pos_t {
	switch input[e.i][e.j] {
	case '.':
		if e.j == 0 && e.dir == "E" {
			return pos_t{e.i, e.j, "E"}
		} else if e.i == 0 && e.dir == "S" {
			return pos_t{e.i, e.j, "S"}
		} else if e.j == len(input[e.i])-1 && e.dir == "W" {
			return pos_t{e.i, e.j, "W"}
		} else if e.i == len(input)-1 && e.dir == "N" {
			return pos_t{e.i, e.j, "N"}
		}
	case '\\':
		if e.j == 0 && e.dir == "E" {
			return pos_t{e.i, e.j, "S"}
		} else if e.i == 0 && e.dir == "S" {
			return pos_t{e.i, e.j, "E"}
		} else if e.j == len(input[e.i])-1 && e.dir == "W" {
			return pos_t{e.i, e.j, "N"}
		} else if e.i == len(input)-1 && e.dir == "N" {
			return pos_t{e.i, e.j, "W"}
		}
	case '-':
		if e.j == 0 && e.dir == "E" {
			return pos_t{e.i, e.j, "E"}
		} else if e.i == 0 && e.dir == "S" {
			if e.j == 0 {
				return pos_t{e.i, e.j, "E"}
			} else {
				return pos_t{e.i, e.j, "W"}
			}
		} else if e.j == len(input[e.i])-1 && e.dir == "W" {
			return pos_t{e.i, e.j, "W"}
		} else if e.i == len(input)-1 && e.dir == "N" {
			if e.j == 0 {
				return pos_t{e.i, e.j, "E"}
			} else {
				return pos_t{e.i, e.j, "W"}
			}
		}
	case '|':
		if e.j == 0 && e.dir == "E" {
			if e.i == 0 {
				return pos_t{e.i, e.j, "S"}
			} else {
				return pos_t{e.i, e.j, "N"}
			}
		} else if e.i == 0 && e.dir == "S" {
			return pos_t{e.i, e.j, "S"}
		} else if e.j == len(input[e.i])-1 && e.dir == "W" {
			if e.i == 0 {
				return pos_t{e.i, e.j, "S"}
			} else {
				return pos_t{e.i, e.j, "N"}
			}
		} else if e.i == len(input)-1 && e.dir == "N" {
			return pos_t{e.i, e.j, "N"}
		}
	default: //case '/':
		if e.j == 0 && e.dir == "E" {
			return pos_t{e.i, e.j, "N"}
		} else if e.i == 0 && e.dir == "S" {
			return pos_t{e.i, e.j, "W"}
		} else if e.j == len(input[e.i])-1 && e.dir == "W" {
			return pos_t{e.i, e.j, "S"}
		} else { //if e.i == len(input)-1 && e.dir == "N" {
			return pos_t{e.i, e.j, "E"}
		}
	}
	return WALL
}

func Part1Parse(input [][]rune, e pos_t) map[pos_t]pos_t {
	m := make(map[pos_t]pos_t)
	var cur_q []pos_t
	var next_q []pos_t
	start := findPos(input, e)
	cur_q = append(cur_q, start)

	for cur_q != nil {
		for _, cur := range cur_q {
			for cur != WALL {
				if _, in_map := m[cur]; in_map {
					break
				} else {
					tm, nexti, nextj := findNext(input, cur)
					m[cur] = tm
					markNext(m, cur, nexti, nextj)
					cur = m[cur]
				}

				if cur.dir == "NS" {
					next_q = append(next_q, pos_t{cur.i, cur.j, "S"})
					cur.dir = "N"
				} else if cur.dir == "WE" {
					next_q = append(next_q, pos_t{cur.i, cur.j, "E"})
					cur.dir = "W"
				}
			}
		}
		cur_q, next_q = next_q, cur_q
		next_q = nil
	}
	return m
}

func findNext(input [][]rune, cur pos_t) (p pos_t, ti, tj int) {
	switch cur.dir {
	case "E":
		j := cur.j + 1
		for ; j < len(input[cur.i]); j++ {
			switch input[cur.i][j] {
			case '|':
				return pos_t{cur.i, j, "NS"}, cur.i, j
			case '/':
				return pos_t{cur.i, j, "N"}, cur.i, j
			case '\\':
				return pos_t{cur.i, j, "S"}, cur.i, j
			}
		}
		tj = j
		ti = cur.i
	case "W":
		j := cur.j - 1
		for ; j >= 0; j-- {
			switch input[cur.i][j] {
			case '|':
				return pos_t{cur.i, j, "NS"}, cur.i, j
			case '/':
				return pos_t{cur.i, j, "S"}, cur.i, j
			case '\\':
				return pos_t{cur.i, j, "N"}, cur.i, j
			}
		}
		tj = j
		ti = cur.i
	case "S":
		i := cur.i + 1
		for ; i < len(input); i++ {
			switch input[i][cur.j] {
			case '-':
				return pos_t{i, cur.j, "WE"}, i, cur.j
			case '/':
				return pos_t{i, cur.j, "W"}, i, cur.j
			case '\\':
				return pos_t{i, cur.j, "E"}, i, cur.j
			}
		}
		ti = i
		tj = cur.j
	case "N":
		i := cur.i - 1
		for ; i >= 0; i-- {
			switch input[i][cur.j] {
			case '-':
				return pos_t{i, cur.j, "WE"}, i, cur.j
			case '/':
				return pos_t{i, cur.j, "E"}, i, cur.j
			case '\\':
				return pos_t{i, cur.j, "W"}, i, cur.j
			}
		}
		ti = i
		tj = cur.j
	}
	return WALL, ti, tj
}

func markNext(m map[pos_t]pos_t, cur pos_t, i, j int) {
	var tmp pos_t
	tmp.dir = cur.dir
	switch cur.dir {
	case "E":
		tmp.i = cur.i
		for tmp.j = cur.j; tmp.j < j; tmp.j++ {
			m[tmp] = m[cur]
		}
	case "W":
		tmp.i = cur.i
		for tmp.j = cur.j; tmp.j > j; tmp.j-- {
			m[tmp] = m[cur]
		}
	case "S":
		tmp.j = cur.j
		for tmp.i = cur.i; tmp.i < i; tmp.i++ {
			m[tmp] = m[cur]
		}
	case "N":
		tmp.j = cur.j
		for tmp.i = cur.i; tmp.i > i; tmp.i-- {
			m[tmp] = m[cur]
		}
	}
}
