package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"strconv"
	"strings"
)

func Part1(input []string) {
	img := Parse1(input)
	start := img.At(500, 0)
	_ = start
	black := img.At(500, 1)
	sand := color.RGBA{240, 230, 140, 0xff}
	end := false
	flag := false
	for !end {
		img.Set(500, 0, sand)
		for j := img.Bounds().Min.Y; j < img.Bounds().Max.Y; j++ {
			for i := img.Bounds().Max.X; i >= img.Bounds().Min.X; i-- {
				if img.At(i, j) == sand {
					if j+1 == img.Bounds().Max.Y {
						img.Set(i, j, black)
						if flag {
							end = true
							break
						}
						if !flag {
							flag = true
							break
						}
					}
					if img.At(i, j+1) == black {
						img.Set(i, j, black)
						img.Set(i, j+1, sand)
					} else if img.At(i-1, j+1) == black {
						img.Set(i, j, black)
						img.Set(i-1, j+1, sand)
					} else if img.At(i+1, j+1) == black {
						img.Set(i, j, black)
						img.Set(i+1, j+1, sand)
					}

				}
			}
		}
	}

	img.Set(500, 0, color.RGBA{255, 0, 0, 0xff})
	sum := 0
	for i := img.Bounds().Min.X; i <= img.Bounds().Max.X; i++ {
		for j := img.Bounds().Min.Y; j <= img.Bounds().Max.Y; j++ {
			if img.At(i, j) == sand {
				sum++
			}
		}
	}
	fmt.Println(sum)
	PrintIMG(img, "part1")
}

func Parse1(input []string) *image.RGBA {
	var lines [][]image.Point
	for _, row := range input {
		fields := strings.Split(row, " -> ")
		var line []image.Point
		for _, f := range fields {
			tmp := strings.Split(f, ",")
			X, err := strconv.Atoi(tmp[0])
			if err != nil {
				log.Fatal(err)
			}
			Y, err := strconv.Atoi(tmp[1])
			if err != nil {
				log.Fatal(err)
			}
			line = append(line, image.Point{X, Y})
		}
		lines = append(lines, line)
	}
	minX := math.MaxInt32
	minY := 0
	maxX, maxY := 0, 0
	for _, l := range lines {
		for _, c := range l {
			if c.X < minX {
				minX = c.X
			}
			if c.Y < minY {
				minY = c.Y
			}
			if c.X > maxX {
				maxX = c.X
			}
			if c.Y > maxY {
				maxY = c.Y
			}
		}
	}
	bounds := image.Rect(minX-1, minY-1, maxX+2, maxY+2)
	img := MapIMG(bounds, lines)
	return img
}
