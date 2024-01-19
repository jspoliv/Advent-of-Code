package main

import "fmt"

type tree_t struct {
	i, j, size int
}

func Part1_2(input [][]rune) {
	grid := Parse1(input)
	sum := 0
	high := 0
	for _, row := range grid {
		for _, g := range row {
			if IsVisible(g, grid) {
				sum++
			}
			if cur := GetScore(g, grid); cur > high {
				high = cur
			}
		}
	}
	fmt.Println(sum)  // part1 result
	fmt.Println(high) // part2 result
}

func Parse1(input [][]rune) [][]tree_t {
	grid := make([][]tree_t, len(input))
	for i := range input {
		grid[i] = make([]tree_t, len(input[i]))
		for j := range input[i] {
			grid[i][j] = tree_t{i: i, j: j, size: int(input[i][j] - '0')}
		}
	}
	return grid
}

func IsVisible(t tree_t, grid [][]tree_t) bool {
	if t.i == 0 || t.j == 0 || t.i == len(grid)-1 || t.j == len(grid[t.i])-1 {
		return true
	}
	u, d, l, r := true, true, true, true
	for i := 0; i < t.i; i++ {
		if grid[i][t.j].size >= t.size {
			u = false
			break
		}
	}
	for i := t.i + 1; i < len(grid); i++ {
		if grid[i][t.j].size >= t.size {
			d = false
			break
		}
	}
	for j := 0; j < t.j; j++ {
		if grid[t.i][j].size >= t.size {
			l = false
			break
		}
	}
	for j := t.j + 1; j < len(grid[t.i]); j++ {
		if grid[t.i][j].size >= t.size {
			r = false
			break
		}
	}
	return l || r || u || d
}

func GetScore(t tree_t, grid [][]tree_t) int {
	if t.i == 0 || t.j == 0 || t.i == len(grid)-1 || t.j == len(grid[t.i])-1 {
		return 0
	}
	u, d, l, r := 0, 0, 0, 0
	for i := t.i - 1; i >= 0; i-- {
		u++
		if grid[i][t.j].size >= t.size {
			break
		}
	}
	for i := t.i + 1; i < len(grid); i++ {
		d++
		if grid[i][t.j].size >= t.size {
			break
		}
	}
	for j := t.j - 1; j >= 0; j-- {
		l++
		if grid[t.i][j].size >= t.size {
			break
		}
	}
	for j := t.j + 1; j < len(grid[t.i]); j++ {
		r++
		if grid[t.i][j].size >= t.size {
			break
		}
	}
	return l * r * u * d
}
