package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type coord struct {
	x, y int
}

func main() {
	file_in, _ := os.Open("sample.txt")
	defer file_in.Close()
	input := parseInput(file_in)
	part1, loop, grid := part1Parse(input)
	fmt.Println(part1)
	fmt.Println(part2(loop, grid))
}

func parseInput(file_in *os.File) []string {
	scanner := bufio.NewScanner(file_in)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func part1Parse(input []string) (i int, loop []coord, grid [][]rune) {
	for _, line := range input {
		grid = append(grid, []rune(line))
	}
	start := coord{-1, -1}
	prox := start
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'S' {
				grid[i][j] = inferS(coord{i, j}, grid, len(grid), len(grid[i]))
				start = coord{i, j}
				if grid[i][j] == '|' {
					prox = coord{i - 1, j}
				} else if grid[i][j] == '-' {
					prox = coord{i, j - 1}
				} else if grid[i][j] == 'L' {
					prox = coord{i - 1, j}
				} else if grid[i][j] == 'J' {
					prox = coord{i - 1, j}
				} else if grid[i][j] == '7' {
					prox = coord{i, j - 1}
				} else { // grid[i][j] == 'F'
					prox = coord{i + 1, j}
				}
				break
			}
		}
		if start.x != -1 {
			break
		}
	}

	prev := start
	loop = append(loop, start)
	cur := prox
	for ; cur.x != start.x || cur.y != start.y; i++ {
		t := cur
		loop = append(loop, t)
		cur = nextPipe(prev, cur, grid)
		prev = t
	}
	//grid[start.x][start.y] = 'S'

	return i/2 + i%2, loop, grid
}

func part2(loop []coord, grid [][]rune) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for l, r := 0, len(grid[i])-1; l <= r; l, r = l+1, r-1 {
			if !coordContains(loop, coord{i, l}) {
				grid[i][l] = '.'
			}
			if !coordContains(loop, coord{i, r}) {
				grid[i][r] = '.'
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == '.' && !wallParity(grid, i, j) {
				count++
			}
		}
	}
	return count
}

func wallParity(grid [][]rune, x, y int) bool {
	var l int
	for j := 0; j < y; j++ {
		val := grid[x][j]
		if val == '|' || val == 'J' || val == 'L' {
			l++
		}
	}
	return l%2 == 0
}

func inferS(c coord, grid [][]rune, lenx, leny int) rune {
	if c.x > 0 && c.x+1 < lenx && strings.ContainsRune("|7F", grid[c.x-1][c.y]) && strings.ContainsRune("|LJ", grid[c.x+1][c.y]) {
		return '|'
	} else if c.y > 0 && c.y+1 < leny && strings.ContainsRune("-LF", grid[c.x][c.y-1]) && strings.ContainsRune("-J7", grid[c.x][c.y+1]) {
		return '-'
	} else if c.x > 0 && c.y+1 < leny && strings.ContainsRune("|7F", grid[c.x-1][c.y]) && strings.ContainsRune("-J7", grid[c.x][c.y+1]) {
		return 'L'
	} else if c.x > 0 && c.y > 0 && strings.ContainsRune("|7F", grid[c.x-1][c.y]) && strings.ContainsRune("-LF", grid[c.x][c.y-1]) {
		return 'J'
	} else if c.x+1 < lenx && c.y > 0 && strings.ContainsRune("|LJ", grid[c.x+1][c.y]) && strings.ContainsRune("-LF", grid[c.x][c.y-1]) {
		return '7'
	} else if c.x+1 < lenx && c.y+1 < leny && strings.ContainsRune("|LJ", grid[c.x+1][c.y]) && strings.ContainsRune("-J7", grid[c.x][c.y+1]) {
		return 'F'
	}
	return 'S'
}

func coordContains(vec []coord, pos coord) bool {
	for _, c := range vec {
		if c.x == pos.x && c.y == pos.y {
			return true
		}
	}
	return false
}

func nextPipe(p, c coord, grid [][]rune) (next coord) {
	switch {
	case grid[p.x][p.y] == '|' && grid[c.x][c.y] == '|' && p.x < c.x:
		next.x = c.x + 1
		next.y = c.y
	case grid[p.x][p.y] == '|' && grid[c.x][c.y] == '|' && p.x > c.x:
		next.x = c.x - 1
		next.y = c.y
	case grid[p.x][p.y] == '|' && (grid[c.x][c.y] == 'J' || grid[c.x][c.y] == '7'):
		next.x = c.x
		next.y = c.y - 1
	case grid[p.x][p.y] == '|' && (grid[c.x][c.y] == 'L' || grid[c.x][c.y] == 'F'):
		next.x = c.x
		next.y = c.y + 1
		//
	case grid[p.x][p.y] == '-' && grid[c.x][c.y] == '-' && p.y < c.y:
		next.x = c.x
		next.y = c.y + 1
	case grid[p.x][p.y] == '-' && grid[c.x][c.y] == '-' && p.y > c.y:
		next.x = c.x
		next.y = c.y - 1
	case grid[p.x][p.y] == '-' && (grid[c.x][c.y] == 'L' || grid[c.x][c.y] == 'J'):
		next.x = c.x - 1
		next.y = c.y
	case grid[p.x][p.y] == '-' && (grid[c.x][c.y] == '7' || grid[c.x][c.y] == 'F'):
		next.x = c.x + 1
		next.y = c.y
		//
	case grid[p.x][p.y] == 'L' && grid[c.x][c.y] == '7' && p.y < c.y:
		next.x = c.x + 1
		next.y = c.y
	case grid[p.x][p.y] == 'L' && grid[c.x][c.y] == '7' && p.x > c.x:
		next.x = c.x
		next.y = c.y - 1
	case grid[p.x][p.y] == 'L' && (grid[c.x][c.y] == 'J' || grid[c.x][c.y] == '|'):
		next.x = c.x - 1
		next.y = c.y
	case grid[p.x][p.y] == 'L' && (grid[c.x][c.y] == 'F' || grid[c.x][c.y] == '-'):
		next.x = c.x
		next.y = c.y + 1
		//
	case grid[p.x][p.y] == 'J' && grid[c.x][c.y] == 'F' && p.x > c.x:
		next.x = c.x
		next.y = c.y + 1
	case grid[p.x][p.y] == 'J' && grid[c.x][c.y] == 'F' && p.y > c.y:
		next.x = c.x + 1
		next.y = c.y
	case grid[p.x][p.y] == 'J' && (grid[c.x][c.y] == 'L' || grid[c.x][c.y] == '|'):
		next.x = c.x - 1
		next.y = c.y
	case grid[p.x][p.y] == 'J' && (grid[c.x][c.y] == '7' || grid[c.x][c.y] == '-'):
		next.x = c.x
		next.y = c.y - 1
		//
	case grid[p.x][p.y] == '7' && grid[c.x][c.y] == 'L' && p.x < c.x:
		next.x = c.x
		next.y = c.y + 1
	case grid[p.x][p.y] == '7' && grid[c.x][c.y] == 'L' && p.y > c.y:
		next.x = c.x - 1
		next.y = c.y
	case grid[p.x][p.y] == '7' && (grid[c.x][c.y] == 'J' || grid[c.x][c.y] == '-'):
		next.x = c.x
		next.y = c.y - 1
	case grid[p.x][p.y] == '7' && (grid[c.x][c.y] == 'F' || grid[c.x][c.y] == '|'):
		next.x = c.x + 1
		next.y = c.y
		//
	case grid[p.x][p.y] == 'F' && grid[c.x][c.y] == 'J' && p.x < c.x:
		next.x = c.x
		next.y = c.y - 1
	case grid[p.x][p.y] == 'F' && grid[c.x][c.y] == 'J' && p.y < c.y:
		next.x = c.x - 1
		next.y = c.y
	case grid[p.x][p.y] == 'F' && (grid[c.x][c.y] == 'L' || grid[c.x][c.y] == '-'):
		next.x = c.x
		next.y = c.y + 1
	case grid[p.x][p.y] == 'F' && (grid[c.x][c.y] == '7' || grid[c.x][c.y] == '|'):
		next.x = c.x + 1
		next.y = c.y
	}
	return
}
