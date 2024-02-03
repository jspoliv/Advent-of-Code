package main

import (
	"fmt"
	"image"
)

func Part2(input []string) {
	pairs, _ := Parse1(input)
	// max := 20 // sample
	max := 4_000_000 // input
	bounds := image.Rect(0, 0, max+1, max+1)
	var c image.Point
	for _, p := range pairs {
		x0 := p.sensor.X - p.dist - 1
		xF := p.sensor.X + p.dist + 1
		for c.X = x0; c.X <= p.sensor.X; c.X++ {
			c.Y = p.sensor.Y - (c.X - x0)
			if c.In(bounds) && check(pairs, c) {
				return
			}
			c.Y += 2 * (c.X - x0)
			if c.In(bounds) && check(pairs, c) {
				return
			}
		}
		for c.X = xF; c.X > p.sensor.X; c.X-- {
			c.Y = p.sensor.Y - (xF - c.X)
			if c.In(bounds) && check(pairs, c) {
				return
			}
			c.Y += 2 * (xF - c.X)
			if c.In(bounds) && check(pairs, c) {
				return
			}
		}
	}
}

func check(pairs []pair_t, c image.Point) bool {
	total := len(pairs)
	found := 0
	for _, p := range pairs {
		if MhtDist(p.sensor, c) > p.dist {
			found++
		} else {
			break
		}
	}
	if found == total {
		fmt.Println(c.X*4000000 + c.Y)
		return true
	}
	return false
}
