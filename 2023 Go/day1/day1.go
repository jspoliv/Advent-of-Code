package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type digit_t struct {
	pos int
	val int
}

func main() {
	file_in, _ := os.Open("input.txt")
	defer file_in.Close()
	parse_digit(file_in)
}

func parse_digit(file_in *os.File) {
	scanner := bufio.NewScanner(file_in)
	file_total := 0
	for scanner.Scan() {
		input := scanner.Text()

		digit := [20]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
			"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
		}

		digit_map := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
			"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
		}

		first := digit_t{999999, -1}
		last := digit_t{-1, -1}
		for _, val := range digit {
			if i := strings.Index(input, val); i != -1 && i < first.pos {
				first.pos = i
				first.val = digit_map[val]
			}

			if i := strings.LastIndex(input, val); i > last.pos {
				last.pos = i
				last.val = digit_map[val]
			}
		}
		file_total += first.val*10 + last.val
	}
	fmt.Println(file_total)
}
