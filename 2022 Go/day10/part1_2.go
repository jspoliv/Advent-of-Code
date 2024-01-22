package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1_2(in []string) {
	sum := 0
	end := 240
	skip := false
	CRT := ""
	for cycle, X, val := 0, 1, 0; cycle < end; {
		for i := 0; i < len(in) && cycle < end; {
			cycle++
			if X-1 == len(CRT) || X == len(CRT) || X+1 == len(CRT) {
				CRT += "#"
			} else {
				CRT += "."
			}
			if len(CRT) == 40 {
				fmt.Println(CRT) // part2 result
				CRT = ""
			}
			if (cycle-20)%40 == 0 {
				sum += X * cycle
			}
			if skip {
				skip = false
				i++
				X += val
				continue
			}
			if in[i][0] == 'a' {
				val, _ = strconv.Atoi(strings.Fields(in[i])[1])
				skip = true
			} else {
				i++
			}
		}
	}
	fmt.Println(sum) // part1 result
}
