package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type set_t struct {
	red   int
	green int
	blue  int
}

type game_t struct {
	id   int
	sets []set_t
}

func main() {
	file_in, _ := os.Open("input.txt")
	defer file_in.Close()
	parse_input(file_in)
}

func parse_input(file_in *os.File) {
	scanner := bufio.NewScanner(file_in)
	game_sum := 0
	power_sum := 0
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ": ") // array with ["Game N", "sets"]
		tmp := strings.Split(input[0], " ")          // array with ["Game", "N"]
		var game game_t
		game.id, _ = strconv.Atoi(tmp[1])   // "N" to int
		tmp = strings.Split(input[1], "; ") // array with separated sets

		for _, a_set := range tmp { // for each set in tmp
			color_set := strings.Split(a_set, ", ") // array with red green blue separated
			tmp_set := set_t{0, 0, 0}
			for _, color := range color_set {
				if strings.Contains(color, "red") {
					tmp3 := strings.Split(color, " ")
					tmp_set.red, _ = strconv.Atoi(tmp3[0])
				}
				if strings.Contains(color, "green") {
					tmp3 := strings.Split(color, " ")
					tmp_set.green, _ = strconv.Atoi(tmp3[0])
				}
				if strings.Contains(color, "blue") {
					tmp3 := strings.Split(color, " ")
					tmp_set.blue, _ = strconv.Atoi(tmp3[0])
				}

			}
			game.sets = append(game.sets, tmp_set)
		}
		flag := true
		tmp_set := set_t{0, 0, 0}
		for _, g := range game.sets {
			if flag && (g.red > 12 || g.green > 13 || g.blue > 14) {
				flag = false
			}
			if g.red > tmp_set.red {
				tmp_set.red = g.red
			}
			if g.green > tmp_set.green {
				tmp_set.green = g.green
			}
			if g.blue > tmp_set.blue {
				tmp_set.blue = g.blue
			}
		}
		power_sum += tmp_set.red * tmp_set.green * tmp_set.blue
		if flag {
			game_sum += game.id
		}
	}
	fmt.Println(game_sum)
	fmt.Println(power_sum)
}
