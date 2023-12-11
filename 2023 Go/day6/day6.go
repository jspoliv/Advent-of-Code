package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type race_t struct {
	time int
	dist int
}

func main() {
	file_in, _ := os.Open("input.txt")
	defer file_in.Close()
	parse_input(file_in)
}

func parse_input(file_in *os.File) {
	scanner := bufio.NewScanner(file_in)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	times := strings.Split(input[0], ": ")
	times[1] = strings.ReplaceAll(times[1], " ", "")
	//times[1] = trim_whitespace(times[1])
	//times = strings.Split(times[1], " ")
	dists := strings.Split(input[1], ": ")
	dists[1] = strings.ReplaceAll(dists[1], " ", "")
	//dists[1] = trim_whitespace(dists[1])
	//dists = strings.Split(dists[1], " ")

	var races []race_t
	//for i := 0; i < len(times); i++ {
	time, _ := strconv.Atoi(times[1])
	dist, _ := strconv.Atoi(dists[1])
	races = append(races, race_t{time, dist})
	//}

	mult := 1
	for _, race := range races {
		sum := 0
		tmp := race.time / 2
		for i := tmp; i > 0; i-- {
			run := (race.time - i) * i
			if run > race.dist {
				sum++
			}
		}
		if race.time%2 == 0 {
			tmp = sum*2 - 1
		} else {
			tmp = sum * 2
		}
		//fmt.Println(tmp)
		mult *= tmp
	}
	fmt.Println(mult)
}

/*func trim_whitespace(s string) string {
	for {
		prev := s
		s = strings.ReplaceAll(s, "  ", " ")
		if s == prev {
			s, _ = strings.CutPrefix(s, " ")
			break
		}
	}
	return s
}*/
