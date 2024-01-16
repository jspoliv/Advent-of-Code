package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type map_t struct {
	out, in, r uint64
}

type seed_t struct {
	seed, r uint64
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

	seed_line := strings.Split(input[0], ": ")
	seed_line = strings.Split(seed_line[1], " ")
	var seeds []seed_t
	/*for _, seed := range seed_line {
		s, _ := strconv.ParseUint(seed, 10, 64)
		seeds = append(seeds, s)
	}*/
	for i := 0; i+1 < len(seed_line); i += 2 {
		var s seed_t
		s.seed, _ = strconv.ParseUint(seed_line[i], 10, 64)
		s.r, _ = strconv.ParseUint(seed_line[i+1], 10, 64)
		seeds = append(seeds, s)
	}

	fmt.Println(seeds)

	i := 1
	soils, i := populate(input, i)
	ferts, i := populate(input, i)
	waters, i := populate(input, i)
	lights, i := populate(input, i)
	temps, i := populate(input, i)
	humids, i := populate(input, i)
	locs, _ := populate(input, i)
	fmt.Println("populated")

	low := uint64(math.MaxUint64)
	for _, seed := range seeds {
		fmt.Printf("seed ")
		for j := uint64(0); j < seed.r; j++ {
			soil := find_map(seed.seed+j, soils)
			fert := find_map(soil, ferts)
			water := find_map(fert, waters)
			light := find_map(water, lights)
			temp := find_map(light, temps)
			humid := find_map(temp, humids)
			loc := find_map(humid, locs)

			if loc < low {
				low = loc
			}
		}
	}
	fmt.Printf("\n%v", low)
}

func find_map(in uint64, m []map_t) uint64 {
	for _, val := range m {
		if val.in <= in && in < val.in+val.r {
			in += val.out - val.in
			break
		}
	}
	return in
}

func populate(input []string, start int) (m []map_t, end int) {
	i := start + 2
	input_len := len(input)
	for ; i < input_len && input[i] != ""; i++ {
		input_line := strings.Split(input[i], " ")
		out, _ := strconv.ParseUint(input_line[0], 10, 64)
		in, _ := strconv.ParseUint(input_line[1], 10, 64)
		rang, _ := strconv.ParseUint(input_line[2], 10, 64)

		m = append(m, map_t{out, in, rang})
	}
	end = i
	return
}
