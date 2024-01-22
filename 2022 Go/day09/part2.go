package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part2(input []string) {
	m2 := Parse2(input)
	fmt.Println(len(m2))
}

func Parse2(input []string) (m2 map[[2]int]struct{}) {
	m2 = make(map[[2]int]struct{})
	var rope [10][2]int
	for _, line := range input {
		t := strings.Fields(line)
		dir := t[0]
		dist, _ := strconv.Atoi(t[1])
		MoveHT9(&rope, dir, dist, m2)
	}
	return
}

func MoveHT9(c *[10][2]int, dir string, dist int, m map[[2]int]struct{}) {
	for d := 0; d < dist; d++ {
		switch dir {
		case "U":
			c[0][0]--
		case "D":
			c[0][0]++
		case "L":
			c[0][1]--
		case "R":
			c[0][1]++
		}
		for i := 1; i < 10; i++ {
			if dX, dY := c[i-1][0]-c[i][0], c[i-1][1]-c[i][1]; math.Abs(float64(dX)) > 1 || math.Abs(float64(dY)) > 1 {
				c[i][0] += n101(dX)
				c[i][1] += n101(dY)
			}
		}
		m[c[9]] = struct{}{}
	}
}

func n101(d int) int {
	if d == 0 {
		return 0
	}
	return int(math.Copysign(1, float64(d)))
}
