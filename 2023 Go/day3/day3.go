package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type star_t struct {
	x, y int
}

type gear_t struct {
	number int
	stars  []star_t
}

func main() {
	input := "input.txt"
	file_in, _ := os.Open(input)
	defer file_in.Close()
	parse_input(file_in)
}

func parse_input(file_in *os.File) {
	scanner := bufio.NewScanner(file_in)
	var matrix [][]rune
	for scanner.Scan() {
		row := []rune(scanner.Text())
		matrix = append(matrix, row)
	}

	part_sum := 0
	number := 0
	valid_part := false
	valid_gear := false
	var gear []gear_t
	var tmp_gear gear_t
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if strings.ContainsRune("0123456789", matrix[i][j]) {
				number = number*10 + int(matrix[i][j]-'0')
				if !valid_part {
					valid_part, valid_gear = valid_neighbors(i, j, matrix, &tmp_gear)
				}
			} else {
				if valid_part {
					if valid_gear {
						tmp_gear.number = number
						gear = append(gear, tmp_gear)
					}
					part_sum += number
					valid_part = false
				}
				number = 0
			}
		}
		if valid_part {
			if valid_gear {
				tmp_gear.number = number
				gear = append(gear, tmp_gear)
			}
			part_sum += number
			valid_part = false
			number = 0
		}
	}
	fmt.Printf("part sum: %v\n", part_sum)

	sum := 0
	for g := 0; g < len(gear)-1; g++ {
		total_adjacent := 0
		pair := -1
		for g2 := g + 1; g2 < len(gear); g2++ {
			for _, s := range gear[g2].stars {
				if slices.Contains(gear[g].stars, s) {
					total_adjacent += 1
					if total_adjacent == 1 {
						pair = g2
					}
				}
			}
		}
		if total_adjacent == 1 {
			//fmt.Printf("g1: %v g2: %v \n", gear[g].number, gear[pair].number)
			sum += gear[g].number * gear[pair].number
		}
	}
	fmt.Printf("gear sum: %v\n", sum)
}

func valid_neighbors(i, j int, matrix [][]rune, tmp_gear *gear_t) (valid_part, valid_gear bool) {
	valid_part = false
	valid_gear = false
	var tmp gear_t
	if i-1 > 0 && j-1 > 0 && !strings.ContainsRune("0123456789.", matrix[i-1][j-1]) {
		valid_part = true
		if matrix[i-1][j-1] == '*' {
			valid_gear = true
			tmp.stars = append(tmp.stars, star_t{i - 1, j - 1})
		}
	}
	if i-1 > 0 && !strings.ContainsRune("0123456789.", matrix[i-1][j]) {
		valid_part = true
		if matrix[i-1][j] == '*' {
			valid_gear = true
			tmp.stars = append(tmp.stars, star_t{i - 1, j})
		}
	}
	if i-1 > 0 && j+1 < len(matrix[i]) && !strings.ContainsRune("0123456789.", matrix[i-1][j+1]) {
		valid_part = true
		if matrix[i-1][j+1] == '*' {
			valid_gear = true
			tmp.stars = append(tmp.stars, star_t{i - 1, j + 1})
		}
	}
	if j-1 > 0 && !strings.ContainsRune("0123456789.", matrix[i][j-1]) {
		valid_part = true
		if matrix[i][j-1] == '*' {
			valid_gear = true
			tmp.stars = append(tmp.stars, star_t{i, j - 1})
		}
	}
	if j+1 < len(matrix[i]) && !strings.ContainsRune("0123456789.", matrix[i][j+1]) {
		valid_part = true
		if matrix[i][j+1] == '*' {
			valid_gear = true
			tmp.stars = append(tmp.stars, star_t{i, j + 1})
		}
	}
	if i+1 < len(matrix) && j-1 > 0 && !strings.ContainsRune("0123456789.", matrix[i+1][j-1]) {
		valid_part = true
		if matrix[i+1][j-1] == '*' {
			valid_gear = true
			tmp.stars = append(tmp.stars, star_t{i + 1, j - 1})
		}
	}
	if i+1 < len(matrix) && !strings.ContainsRune("0123456789.", matrix[i+1][j]) {
		valid_part = true
		if matrix[i+1][j] == '*' {
			valid_gear = true
			tmp.stars = append(tmp.stars, star_t{i + 1, j})
		}
	}
	if i+1 < len(matrix) && j+1 < len(matrix[i]) && !strings.ContainsRune("0123456789.", matrix[i+1][j+1]) {
		valid_part = true
		if matrix[i+1][j+1] == '*' {
			valid_gear = true
			tmp.stars = append(tmp.stars, star_t{i + 1, j + 1})
		}
	}
	if valid_gear {
		*tmp_gear = tmp
	}
	return
}
