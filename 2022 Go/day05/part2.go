package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Move maintaining order
func Part2(input []string) {
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
		s[t5] += s[t3][len(s[t3])-t1:]
		s[t3] = s[t3][:len(s[t3])-t1]
	}
	sum := ""
	for _, e := range s {
		if len(e) > 0 {
			sum += string(e[len(e)-1])
		}
	}
	fmt.Println(sum)
}
