package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type lens_t struct {
	label string
	val   int
}

func main() {
	file := ".txt"
	if len(os.Args) == 1 {
		file = "sample" + file
	} else {
		file = os.Args[1] + file
	}
	file_in, _ := os.Open(file)
	fmt.Println(file)
	defer file_in.Close()
	input := ParseInput(file_in)
	//Part1(input)
	Part2(input)
}

func ParseInput(file_in *os.File) (input []string) {
	scanner := bufio.NewScanner(file_in)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func Part1Parse(input []string) []string {
	in := strings.Split(input[0], ",")
	return in
}

func Part2Parse(input []string) (box [256][]lens_t) {
	inp := Part1Parse(input)

	for _, e := range inp {
		if e[len(e)-1] == '-' {
			dash_lb, _ := strings.CutSuffix(e, "-")
			lb_hash := getHash(dash_lb)
			if j := BoxContains(box, lb_hash, dash_lb); j != -1 {
				box[lb_hash] = append(box[lb_hash][:j], box[lb_hash][j+1:]...)
			}
		} else {
			label := strings.Split(e, "=")
			lb_hash := getHash(label[0])
			lens, _ := strconv.Atoi(label[1])

			if j := BoxContains(box, lb_hash, label[0]); j != -1 {
				box[lb_hash][j].val = lens
			} else {
				box[lb_hash] = append(box[lb_hash], lens_t{label[0], lens})
			}
		}
	}
	return
}

func Part2(input []string) {
	box := Part2Parse(input)
	sum := 0
	for i := 0; i < len(box); i++ {
		if box[i] != nil {
			for j := 0; j < len(box[i]); j++ {
				sum += (i + 1) * (j + 1) * box[i][j].val
			}
		}
	}
	fmt.Println(sum)
}

func Part1(input []string) {
	in := Part1Parse(input)
	sum := 0
	for _, step := range in {
		cur := getHash(step)
		sum += cur
	}
	fmt.Println(sum)
}

func BoxContains(box [256][]lens_t, hash int, label string) (pos int) {
	for j, b := range box[hash] {
		if b.label == label {
			return j
		}
	}
	return -1
}

func getHash(step string) int {
	cur := 0
	for i := range step {
		cur += int(step[i])
		cur = cur * 17
		cur = cur % 256
	}
	return cur
}
