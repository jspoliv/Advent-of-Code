package main

import "fmt"

func Part2(input []string) {
	pair := Parse1(input)
	sum := 0
	for _, p := range pair {
		if betw(p[0].s, p[1].s, p[1].e) || betw(p[0].e, p[1].s, p[1].e) || betw(p[1].s, p[0].s, p[0].e) || betw(p[1].e, p[0].s, p[0].e) {
			sum++
		}
	}
	fmt.Println(sum)
}

// if ( a <= val <= b )
func betw(val, a, b int) bool { return a <= val && val <= b }
