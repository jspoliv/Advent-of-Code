package main

import (
	"fmt"
	"strconv"
	"strings"
)

type range_t struct {
	s, e int // start, end
}

func Part1(input []string) {
	pair := Parse1(input)
	sum := 0
	for _, p := range pair {
		if (p[0].s >= p[1].s && p[0].e <= p[1].e) || (p[1].s >= p[0].s && p[1].e <= p[0].e) {
			sum++
		}
	}
	fmt.Println(sum)
}

func Parse1(input []string) [][2]range_t {
	var pair [][2]range_t
	for _, line := range input {
		t := strings.Split(line, ",")
		left := strings.Split(t[0], "-")
		right := strings.Split(t[1], "-")
		var l, r range_t
		l.s, _ = strconv.Atoi(left[0])
		l.e, _ = strconv.Atoi(left[1])
		r.s, _ = strconv.Atoi(right[0])
		r.e, _ = strconv.Atoi(right[1])
		pair = append(pair, [2]range_t{l, r})
	}
	return pair
}
