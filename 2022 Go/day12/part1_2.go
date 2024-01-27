package main

import (
	"fmt"
	"image"
)

func Part1_2(input [][]rune) {
	start, end := Parse1(input)
	fmt.Println(Dijkstra(start, end, input)) // result part1
	low := INFINITY
	for r := range input {
		for c := range input[r] {
			if input[r][c] == 'a' {
				if d := Dijkstra(image.Point{r, c}, end, input); d < low && d != -1 {
					low = d
				}
			}
		}
	}
	fmt.Println(low) // result part2
}

func Parse1(input [][]rune) (start, end image.Point) {
	for r := range input {
		for c := range input[r] {
			if input[r][c] == 'S' {
				start.X = r
				start.Y = c
				input[r][c] = 'a'
			} else if input[r][c] == 'E' {
				end.X = r
				end.Y = c
				input[r][c] = 'z'
			}
		}
	}
	return
}
