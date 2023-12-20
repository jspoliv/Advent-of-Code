package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	file := "sample"
	file_in, _ := os.Open(file + ".txt")
	fmt.Println(file + ".txt")
	defer file_in.Close()
	input := parseInput(file_in)
	part1(input)
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

func part1Parse(input []string) (pattern [][]string) {
	var pat []string
	for _, row := range input {
		if row == "" {
			pattern = append(pattern, pat)
			pat = nil
			continue
		}
		pat = append(pat, row)
	}
	pattern = append(pattern, pat)
	return
}

func part1(input []string) {
	sum := 0
	for _, p := range part1Parse(input) {
		val := check1(transpStr(p))
		if val != -1 {
			sum += val
		} else {
			sum += 100 * check1(p)
		}
	}
	fmt.Println(sum)
}

func part2(input []string) {
	sum := 0
	for _, p := range part1Parse(input) {
		val := check2(transpStr(p))
		if val != -1 {
			sum += val
		} else {
			sum += 100 * check2(p)
		}
	}
	fmt.Println(sum)
}

func check2(pat []string) (pos int) {
	for i := 0; i < len(pat)-1; i++ {
		dif := diff(pat[i], pat[i+1])
		smudge := false
		if dif == 1 {
			smudge = true
		}
		if dif == 0 || dif == 1 {
			r := i + 2
			l := 2*i + 1 - r
			for 0 <= l && r < len(pat) {
				dif = diff(pat[l], pat[r])
				if dif > 1 || (dif == 1 && smudge) {
					break
				} else if dif == 1 {
					smudge = true
				}
				r++
				l--
			}
			if !(0 <= l && r < len(pat)) {
				if smudge {
					return i + 1
				}
			}
		}
	}
	return -1
}

func check1(pat []string) (pos int) {
	for i := 0; i < len(pat)-1; i++ {
		if pat[i] == pat[i+1] {
			r := i + 2
			l := 2*i + 1 - r
			for 0 <= l && r < len(pat) {
				if pat[l] != pat[r] {
					break
				}
				r++
				l--
			}
			if !(0 <= l && r < len(pat)) {
				return i + 1
			}
		}
	}
	return -1
}

func diff(a, b string) (count int) {
	if a != b {
		t := a
		if len(a) > len(b) {
			t = b
		}
		for i := range t {
			if a[i] != b[i] {
				count++
			}
		}
	}
	return count + int(math.Abs(float64(len(a)-len(b))))
}

func transpStr(input []string) []string {
	transp := make([][]rune, len(input[0]))
	for _, row := range input {
		for j, char := range row {
			transp[j] = append(transp[j], char)
		}
	}
	var tmp []string
	for _, row := range transp {
		tmp = append(tmp, string(row))
	}
	return tmp
}
