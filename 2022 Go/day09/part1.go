package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part1(input []string) {
	m1 := Parse1(input)
	fmt.Println(len(m1))
}

func Parse1(input []string) (m1 map[[2]int]struct{}) {
	m1 = make(map[[2]int]struct{})
	var h1, t1 [2]int
	for _, line := range input {
		t := strings.Fields(line)
		dir := t[0]
		dist, _ := strconv.Atoi(t[1])
		MoveHT1(&h1, &t1, dir, dist, m1)
	}
	return
}

func HTdist(h, t [2]int) int {
	if h[0] == t[0] {
		if math.Abs(float64(h[1]-t[1])) <= 1 {
			return 0
		}
	} else if h[1] == t[1] {
		if math.Abs(float64(h[0]-t[0])) <= 1 {
			return 0
		}
	} else if math.Abs(float64(h[0]-t[0])) == 1 && math.Abs(float64(h[1]-t[1])) == 1 {
		return 0
	}
	return 1
}

func MoveHT1(h, t *[2]int, dir string, dist int, m map[[2]int]struct{}) {
	curh, curt := *h, *t
	for i := 0; i < dist; i++ {
		prevh := curh
		switch dir {
		case "U":
			curh[0]--
		case "D":
			curh[0]++
		case "L":
			curh[1]--
		case "R":
			curh[1]++
		}
		if HTdist(curh, curt) > 0 {
			curt = prevh
		}
		m[curt] = struct{}{}
	}
	*h = curh
	*t = curt
}
