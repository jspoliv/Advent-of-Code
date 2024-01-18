package main

import "fmt"

const P2 int = 14

func Part2(input []string) {
	diff := P2
	s := input[0]
	result := -1
	seen := make(map[byte]struct{})
	for i := diff - 1; i < len(s); i++ {
		for j := 0; j < diff; j++ {
			seen[s[i-j]] = struct{}{}
		}
		if len(seen) == diff {
			result = i + 1
			break
		}
		for j := 0; j < diff; j++ {
			delete(seen, s[i-j])
		}
	}
	fmt.Println(result)
}
