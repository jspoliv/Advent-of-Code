package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	file := "input"
	file_in, _ := os.Open(file + ".txt")
	fmt.Println(file + ".txt")
	defer file_in.Close()
	input := parseInput(file_in)
	//Part1(input)
	Part2(input)
}

func parseInput(file_in *os.File) [][]rune {
	scanner := bufio.NewScanner(file_in)
	var input [][]rune
	for scanner.Scan() {
		input = append(input, []rune(scanner.Text()))
	}
	return input
}

func part1Parse(input [][]rune) [][]rune {
	p := input
	tiltNorth(p)
	p = rotate(p)
	p = rotate(p)
	return p
}

func Part1(input [][]rune) {
	fmt.Println("part1")
	p := part1Parse(input)
	sum := 0
	for i := range p {
		sum += strings.Count(string(p[i]), "O") * (i + 1)
	}
	fmt.Println(sum)
}

func SpinCycle(p [][]rune) {
	tiltNorth(p)
	tiltWest(p)
	tiltSouth(p)
	tiltEast(p)
}

func part2Parse(input [][]rune) [][]rune {
	p := input
	spins := 118 // 118 == 1bil for input.txt
	/*
		spins := 1_000_000_000
		m := make(map[string]int)
		cycle_s := -1
		cycle_e := -1
		for i := 0; i < spins; i++ {
			SpinCycle(p)
			tmp := ToString(p)
			e, b := m[tmp]
			if b {
				cycle_s = e
				cycle_e = i
				break
			} else {
				m[tmp] = i
			}
		}
		fmt.Printf("Cycle: %v -> %v\n", cycle_s, cycle_e)
		total_spins := cycle_s + (spins-cycle_s)%(cycle_e-cycle_s)
		fmt.Printf("1bil == %v spins\n", total_spins)
	*/
	for i := 0; i < spins; i++ {
		SpinCycle(p)
	}
	p = rotate(p)
	p = rotate(p)
	return p
}

func ToString(in [][]rune) string {
	out := ""
	for _, r := range in {
		out += string(r)
	}
	return out
}

func Part2(input [][]rune) {
	fmt.Println("part2")
	p := part2Parse(input)
	sum := 0
	for i := range p {
		sum += strings.Count(string(p[i]), "O") * (i + 1)
	}
	// sum is 89845 for input.txt
	fmt.Println(sum)
}

func rotate(in [][]rune) [][]rune {
	len_in := len(in)
	len_ini := len(in[0])
	out := make([][]rune, len_in)
	var wg sync.WaitGroup
	for i := 0; i < len_in; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			out[i] = make([]rune, len_ini)
		}()
	}
	wg.Wait()
	for i := 0; i < len_in; i++ {
		for j := 0; j < len_ini; j++ {
			wg.Add(1)
			i := i
			j := j
			go func() {
				defer wg.Done()
				out[i][j] = in[len_in-1-j][i]
			}()
		}
	}
	wg.Wait()
	return out
}

func tiltSouth(g [][]rune) {
	var wg sync.WaitGroup
	for i := range g {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			len_s := len(g)
			count := countOy(g, i)
			for c := 0; c < count; c++ {
				for j := 0; j < len_s-1; j++ {
					if g[j][i] == 'O' && g[j+1][i] == '.' {
						g[j][i], g[j+1][i] = g[j+1][i], g[j][i]
					}
				}
			}
		}()
	}
	wg.Wait()
}

func tiltNorth(g [][]rune) {
	var wg sync.WaitGroup
	for i := range g {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			len_s := len(g)
			count := countOy(g, i)
			for c := 0; c < count; c++ {
				for j := len_s - 1; j > 0; j-- {
					if g[j][i] == 'O' && g[j-1][i] == '.' {
						g[j][i], g[j-1][i] = g[j-1][i], g[j][i]
					}
				}
			}
		}()
	}
	wg.Wait()
}

func tiltEast(g [][]rune) {
	var wg sync.WaitGroup
	for i := range g {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			len_s := len(g[i])
			count := countOx(g, i)
			for c := 0; c < count; c++ {
				for j := 0; j < len_s-1; j++ {
					if g[i][j] == 'O' && g[i][j+1] == '.' {
						g[i][j], g[i][j+1] = g[i][j+1], g[i][j]
					}
				}
			}
		}()
	}
	wg.Wait()
}

func tiltWest(g [][]rune) {
	var wg sync.WaitGroup
	for i := range g {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			len_s := len(g[i])
			count := countOx(g, i)
			for c := 0; c < count; c++ {
				for j := len_s - 1; j > 0; j-- {
					if g[i][j] == 'O' && g[i][j-1] == '.' {
						g[i][j], g[i][j-1] = g[i][j-1], g[i][j]
					}
				}
			}
		}()
	}
	wg.Wait()
}

func countOy(g [][]rune, j int) int {
	len_i := len(g)
	count := 0
	for i := 0; i < len_i; i++ {
		if g[i][j] == 'O' {
			count++
		}
	}
	return count
}

func countOx(g [][]rune, i int) int {
	len_j := len(g[i])
	count := 0
	for j := 0; j < len_j; j++ {
		if g[i][j] == 'O' {
			count++
		}
	}
	return count
}
