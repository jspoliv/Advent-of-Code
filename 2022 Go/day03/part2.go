package main

import (
	"fmt"
	"strings"
)

func Part2(input []string) {
	r := input
	sum := 0
	for i := 2; i < len(r); i += 3 {
		seen := make(map[byte]struct{})
		for j := range r[i] {
			if _, f := seen[r[i][j]]; !f && strings.Contains(r[i-1], string(r[i][j])) && strings.Contains(r[i-2], string(r[i][j])) {
				seen[r[i][j]] = struct{}{}
				sum += Rtoi(r[i][j])
			}
		}
	}
	fmt.Println(sum)
}
