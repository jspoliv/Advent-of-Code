package main

import (
	"fmt"
	"image"
	"log"
	"math"
	"strconv"
	"strings"
)

type pair_t struct {
	sensor, beacon image.Point
	dist           int
}

func Part1(input []string) {
	pairs, bounds := Parse1(input)
	// y := 10      // sample max
	y := 2000000 // input max
	beacon := make(map[image.Point]struct{})
	inrange := make(map[image.Point]struct{})
	cur := image.Point{Y: y}
	for cur.X = bounds.Min.X; cur.X <= bounds.Max.X; cur.X++ {
		for _, p := range pairs {
			if p.dist >= MhtDist(p.sensor, cur) {
				if cur == p.beacon {
					beacon[cur] = struct{}{}
				}
				inrange[cur] = struct{}{}
			}
		}
	}
	for k := range beacon {
		delete(inrange, k)
	}
	fmt.Println(len(inrange))
}

func Parse1(input []string) (pairs []pair_t, bounds image.Rectangle) {
	maxX, minX := math.MinInt32, math.MaxInt32
	maxY, minY := math.MinInt32, math.MaxInt32
	for _, s := range input {
		s, _ := strings.CutPrefix(s, "Sensor at ")
		t1 := strings.Split(s, ": closest beacon is at ")
		t2 := strings.Split(t1[0], ", ")
		tx, err := strconv.Atoi(strings.Split(t2[0], "=")[1])
		if err != nil {
			log.Fatal(err)
		}
		ty, err := strconv.Atoi(strings.Split(t2[1], "=")[1])
		if err != nil {
			log.Fatal(err)
		}
		sensor := image.Point{tx, ty}
		t2 = strings.Split(t1[1], ", ")
		tx, err = strconv.Atoi(strings.Split(t2[0], "=")[1])
		if err != nil {
			log.Fatal(err)
		}
		ty, err = strconv.Atoi(strings.Split(t2[1], "=")[1])
		if err != nil {
			log.Fatal(err)
		}
		beacon := image.Point{tx, ty}
		if d := MhtDist(sensor, beacon); sensor.X+d > maxX {
			maxX = sensor.X + d
		} else if d := MhtDist(sensor, beacon); sensor.X-d < minX {
			minX = sensor.X - d
		}
		if d := MhtDist(sensor, beacon); sensor.Y+d > maxY {
			maxY = sensor.Y + d
		} else if d := MhtDist(sensor, beacon); sensor.Y-d < minY {
			minY = sensor.Y - d
		}
		pairs = append(pairs, pair_t{sensor: sensor, beacon: beacon, dist: MhtDist(sensor, beacon)})
	}
	bounds = image.Rect(minX, minY, maxX+1, maxY+1)
	return pairs, bounds
}

func MhtDist(A, B image.Point) (r int) {
	return int(math.Abs(float64(A.X-B.X)) + math.Abs(float64(A.Y-B.Y)))
}
