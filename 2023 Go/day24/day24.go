package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := ".txt"
	if len(os.Args) == 1 {
		file = "sample" + file
	} else { // go run . [filepath]
		file = os.Args[1] + file
	}
	fmt.Println(file)
	input := ParseInput(file)
	Part2(input)
}

func ParseInput(filepath string) (input []string) {
	file_in, _ := os.Open(filepath)
	defer file_in.Close()
	scanner := bufio.NewScanner(file_in)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func Part1Parse(input []string) (h_list []hail_t) {
	for _, line := range input {
		pos := strings.ReplaceAll(line, ",", "")
		tmp := strings.Split(pos, "@")
		pos = tmp[0]
		vel := tmp[1]

		var hail hail_t
		tmp = strings.Fields(pos)
		hail.p.x, _ = strconv.ParseFloat(tmp[0], 64)
		hail.p.y, _ = strconv.ParseFloat(tmp[1], 64)
		hail.p.z, _ = strconv.ParseFloat(tmp[2], 64)
		tmp = strings.Fields(vel)
		hail.v.x, _ = strconv.ParseFloat(tmp[0], 64)
		hail.v.y, _ = strconv.ParseFloat(tmp[1], 64)
		hail.v.z, _ = strconv.ParseFloat(tmp[2], 64)

		h_list = append(h_list, hail)
	}
	return
}

type hail_t struct {
	p, v vector_t
}

type vector_t struct {
	x, y, z float64
}

// in (7,7) (27,27)
const (
	LOW_END  float64 = 200000000000000
	HIGH_END float64 = 400000000000000
	CENTER   float64 = (HIGH_END-LOW_END)/2 + LOW_END
)

func (v vector_t) In() bool {
	return v.x >= LOW_END && v.x <= HIGH_END && v.y >= LOW_END && v.y <= HIGH_END
}
func (a vector_t) Add(b vector_t, times ...float64) vector_t {
	if len(times) == 1 {
		return vector_t{a.x + b.x*times[0], a.y + b.y*times[0], a.z + b.z*times[0]}
	} else {
		return vector_t{a.x + b.x, a.y + b.y, a.z + b.z}
	}
}
func (a vector_t) Sub(b vector_t, times ...float64) vector_t {
	if len(times) == 1 {
		return vector_t{a.x - b.x*times[0], a.y - b.y*times[0], a.z - b.z*times[0]}
	} else {
		return vector_t{a.x - b.x, a.y - b.y, a.z - b.z}
	}
}
func (v vector_t) String() string {
	return fmt.Sprintf("(%v,%v)", v.x, v.y)
}
func (h hail_t) String() string {
	return fmt.Sprintf("%.0f, %.0f, %.0f @ %4.0f, %4.0f, %4.0f", h.p.x, h.p.y, h.p.z, h.v.x, h.v.y, h.v.z)
}

func Part1(input []string) {
	h := Part1Parse(input)
	//fmt.Println(h)
	sum := 0
	for a := 0; a < len(h)-1; a++ {
		for b := a + 1; b < len(h); b++ {
			sum += Colides(h[a], h[b])
		}
	}
	fmt.Println(sum)
}

func Colides(a, b hail_t) int {
	a1, b1 := a.p, b.p
	a2, b2 := a1.Add(a.v), b1.Add(b.v)
	x1, x2, x3, x4 := a1.x, a2.x, b1.x, b2.x
	y1, y2, y3, y4 := a1.y, a2.y, b1.y, b2.y

	// lines of lenght zero
	if (x1 == x2 && y1 == y2) || (x3 == x4 && y3 == y4) {
		return 0
	}
	denom := (y4-y3)*(x2-x1) - (x4-x3)*(y2-y1)
	// lines are parallel
	if denom == 0 {
		return 0
	}
	ua := ((x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)) / denom
	ub := ((x2-x1)*(y1-y3) - (y2-y1)*(x1-x3)) / denom

	x := x1 + ua*(x2-x1)
	y := y1 + ua*(y2-y1)
	p := vector_t{x, y, 0}

	if ua <= 0 || ub <= 0 {
		return 0
	}

	if !p.In() {
		return 0
	}

	return 1
}

func Part2Parse(input []string) (h_list []hail_t) {
	for i, line := range input {
		if i == 3 {
			break
		}
		pos := strings.ReplaceAll(line, ",", "")
		tmp := strings.Split(pos, "@")
		pos = tmp[0]
		vel := tmp[1]

		var hail hail_t
		tmp = strings.Fields(pos)
		hail.p.x, _ = strconv.ParseFloat(tmp[0], 64)
		hail.p.y, _ = strconv.ParseFloat(tmp[1], 64)
		hail.p.z, _ = strconv.ParseFloat(tmp[2], 64)
		tmp = strings.Fields(vel)
		hail.v.x, _ = strconv.ParseFloat(tmp[0], 64)
		hail.v.y, _ = strconv.ParseFloat(tmp[1], 64)
		hail.v.z, _ = strconv.ParseFloat(tmp[2], 64)

		h_list = append(h_list, hail)
	}
	return
}

func Part2(input []string) {
	hlist := Part2Parse(input)

	for _, h := range hlist {
		fmt.Println(h)
	}
}
