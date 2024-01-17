package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Push(s string, b byte) string {
	return s + string(b)
}
func Pop(s string) (s1 string, b byte) {
	b = s[len(s)-1]
	s1 = s[:len(s)-1]
	return s1, b
}

// Move and inverse order
func Part1(input []string) {
	s_len := (len(input[0])-3)/4 + 1
	s := make([]string, s_len)
	brk := 0
	// parse stacks
	for b, row := range input {
		if !strings.Contains(row, "[") {
			brk = b + 2
			break
		}
		for i, j := 1, 0; i < len(row); i += 4 {
			if row[i] != ' ' {
				s[j] = string(row[i]) + s[j]
			}
			j++
		}
	}
	// parse movement
	for i := brk; i < len(input); i++ {
		t := strings.Fields(input[i])
		t1, _ := strconv.Atoi(t[1]) // move t1 elements
		t3, _ := strconv.Atoi(t[3]) // from pos t3
		t5, _ := strconv.Atoi(t[5]) // to pos t5
		t3, t5 = t3-1, t5-1
		for j := 0; j < t1; j++ {
			ts, by := Pop(s[t3])
			s[t3] = ts
			s[t5] = Push(s[t5], by)
		}
	}
	sum := ""
	for _, e := range s {
		sum += string(e[len(e)-1])
	}
	fmt.Println(sum)
}
