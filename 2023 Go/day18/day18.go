package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := ".txt"
	if len(os.Args) == 1 {
		file = "sample" + file
	} else {
		file = os.Args[1] + file
	}
	fmt.Println(file)
	file_in, _ := os.Open(file)
	defer file_in.Close()
	input := ParseInput(file_in)
	Part1(input)
	Part2(input)
}

func ParseInput(file_in *os.File) (input []string) {
	scanner := bufio.NewScanner(file_in)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

type step_t struct {
	dir  string
	dist int
}

func Part2Parse(input []string) (dig_plan []step_t) {
	defer fmt.Println("parse2 done")
	for _, row := range input {
		thex := strings.Fields(row)[2]
		thex = thex[2 : len(thex)-1]
		tdir := thex[len(thex)-1:]
		thex = thex[:len(thex)-1]
		tdist, _ := strconv.ParseUint(thex, 16, 64)

		switch tdir {
		case "0":
			tdir = "R"
		case "1":
			tdir = "D"
		case "2":
			tdir = "L"
		case "3":
			tdir = "U"
		}
		dig_plan = append(dig_plan, step_t{tdir, int(tdist)})
	}
	return dig_plan
}

// Part2 works for Part1 as well by swapping Part2Parse with Part1Parse
func Part2(input []string) {
	dig_plan := Part2Parse(input)
	dimX, dimY, limX, limY := findDim(dig_plan)
	start := [2]int{dimX - 1 - limX, dimY - 1 - limY}
	fmt.Printf("Size: %v x %v\nStart: (%v,%v)\n", dimX, dimY, start[0], start[1])

	var point [][2]int
	i := start[0]
	j := start[1]
	point = append(point, [2]int{i, j})
	perimeter := 0
	for _, cur := range dig_plan {
		switch cur.dir {
		case "U":
			i -= cur.dist
		case "D":
			i += cur.dist
		case "L":
			j -= cur.dist
		case "R":
			j += cur.dist
		}
		point = append(point, [2]int{i, j})
		perimeter += cur.dist
	}

	area := 0
	for t := 1; t < len(point); t++ {
		area += Det2(point[t-1], point[t])
	}
	area = area / 2
	area = int(math.Abs(float64(area))) // shoelace formula
	fmt.Println(area + perimeter/2 + 1) // pick's theorem
}

func Det2(A, B [2]int) int { return A[0]*B[1] - B[0]*A[1] }

func Part1Parse(input []string) (dig_plan []step_t) {
	defer fmt.Println("parse1 done")
	for _, row := range input {
		tfields := strings.Fields(row)
		tdir := tfields[0]
		tdist, _ := strconv.Atoi(tfields[1])
		dig_plan = append(dig_plan, step_t{tdir, tdist})
	}
	return dig_plan
}

func Part1(input []string) {
	dig_plan := Part1Parse(input)
	dimX, dimY, limX, limY := findDim(dig_plan)
	plan := mapPlan(dig_plan, dimX, dimY, limX, limY)

	tmp := make([]rune, len(plan[0]))
	for i := 0; i < len(tmp); i++ {
		tmp[i] = '.'
	}
	plan = append([][]rune{tmp}, plan...)
	plan = append(plan, tmp)
	for i := 0; i < len(plan); i++ {
		plan[i] = append([]rune{'.'}, plan[i]...)
		plan[i] = append(plan[i], '.')
	}
	fmt.Println("append border done")

	var queue [][2]int
	var next [][2]int
	plan[0][0] = 'O'
	queue = append(queue, [2]int{0, 0})

	for queue != nil {
		for _, cur := range queue {
			next = append(next, Fill(plan, cur)...)
		}
		queue, next = next, queue
		next = nil
	}
	fmt.Println("fill done")

	plan = plan[1 : len(plan)-1]
	for i := 0; i < len(plan); i++ {
		plan[i] = plan[i][1 : len(plan[i])-1]
	}
	fmt.Println("trim border done")

	count := 0
	for _, row := range plan {
		for _, r := range row {
			if r != 'O' {
				count++
			}
		}
	}
	fmt.Println(count)
}

func mapPlan(dig_plan []step_t, dimX, dimY, limX, limY int) [][]rune {
	defer fmt.Println("mapPlan() done")
	plan := make([][]rune, dimX)
	for i := 0; i < dimX; i++ {
		plan[i] = make([]rune, dimY)
		for j := 0; j < len(plan[i]); j++ {
			plan[i][j] = '.'
		}
	}

	i, j := len(plan)-1-limX, len(plan[0])-1-limY
	for _, s := range dig_plan {
		switch s.dir {
		case "L":
			for l := 0; l < s.dist; l++ {
				j--
				plan[i][j] = '#'
			}
		case "R":
			for r := 0; r < s.dist; r++ {
				j++
				plan[i][j] = '#'
			}
		case "U":
			for u := 0; u < s.dist; u++ {
				i--
				plan[i][j] = '#'
			}
		default: // case "D"
			for d := 0; d < s.dist; d++ {
				i++
				plan[i][j] = '#'
			}
		}
	}
	return plan
}

func Fill(p [][]rune, cur [2]int) (neighbors [][2]int) {
	i, j := cur[0], cur[1]
	if p[i][j] == 'O' {
		if i-1 > 0 && p[i-1][j] != '#' && p[i-1][j] != 'O' {
			p[i-1][j] = 'O'
			neighbors = append(neighbors, [2]int{i - 1, j})
		}
		if j-1 > 0 && p[i][j-1] != '#' && p[i][j-1] != 'O' {
			p[i][j-1] = 'O'
			neighbors = append(neighbors, [2]int{i, j - 1})
		}
		if i+1 < len(p) && p[i+1][j] != '#' && p[i+1][j] != 'O' {
			p[i+1][j] = 'O'
			neighbors = append(neighbors, [2]int{i + 1, j})
		}
		if j+1 < len(p[i]) && p[i][j+1] != '#' && p[i][j+1] != 'O' {
			p[i][j+1] = 'O'
			neighbors = append(neighbors, [2]int{i, j + 1})
		}
	}
	return
}

func findDim(dig_plan []step_t) (dimX, dimY, limX, limY int) {
	sumX, sumY, lowX, lowY := 0, 0, 0, 0
	for _, step := range dig_plan {
		switch step.dir {
		case "L":
			sumY -= step.dist
			if sumY < lowY {
				lowY = sumY
			}
		case "R":
			sumY += step.dist
			if sumY > dimY {
				dimY = sumY
			}
		case "U":
			sumX -= step.dist
			if sumX < lowX {
				lowX = sumX
			}
		default: // case "D"
			sumX += step.dist
			if sumX > dimX {
				dimX = sumX
			}
		}
	}
	fmt.Printf("X:[%v..%v] Y:[%v..%v]\n", lowX, dimX, lowY, dimY)
	return dimX + 1 - lowX, dimY + 1 - lowY, dimX, dimY
}
