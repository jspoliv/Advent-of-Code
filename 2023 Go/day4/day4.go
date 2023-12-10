package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type card_t struct {
	repeats int
	info    string
}

func main() {
	file_in, _ := os.Open("input.txt")
	defer file_in.Close()
	parse_input(file_in)
}

func parse_input(file_in *os.File) {
	scanner := bufio.NewScanner(file_in)
	var deck []card_t
	for scanner.Scan() {
		deck = append(deck, card_t{1, scanner.Text()})
	}
	//total_score := 0
	card_amount := 0
	for c := range deck {
		card_numbers, win_numbers := parse_info(deck[c].info)
		win_amount := check_wins(card_numbers, win_numbers)
		for i := 0; i < deck[c].repeats; i++ {
			card_amount += 1
			if win_amount > 0 {
				//total_score += int(math.Pow(2, float64(win_amount-1)))
				for j := 1; j <= win_amount && c+j < len(deck); j++ {
					deck[c+j].repeats += 1
				}
			}
		}
	}
	fmt.Println(card_amount)
}

func parse_info(info string) (card_numbers, win_numbers []string) {
	card_numbers = strings.Split(info, ": ")
	win_numbers = strings.Split(card_numbers[1], " | ")
	card_numbers = strings.Split(win_numbers[1], " ")
	win_numbers = strings.Split(win_numbers[0], " ")
	return
}

func check_wins(card_numbers, win_numbers []string) (win_amount int) {
	win_amount = 0
	for _, win := range win_numbers {
		if win == "" {
			continue
		}
		for _, number := range card_numbers {
			if win == number {
				win_amount += 1
			}
		}
	}
	return
}
