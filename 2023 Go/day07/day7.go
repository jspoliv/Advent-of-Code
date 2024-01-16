package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand_t struct {
	cards    string
	bid      int
	type_h   string
	strength string
}

var j1_map = map[string]string{"five": "five", "four": "five", "full house": "four", "three": "four",
	"two pair": "full house", "one pair": "three", "high card": "one pair"}

var j2_map = map[string]string{"full house": "five", "two pair": "four", "one pair": "three"}

var j3_map = map[string]string{"full house": "five", "three": "four"}

var strength_map = map[rune]string{'A': "m", 'K': "l", 'Q': "k", 'J': "a", 'T': "j", '9': "i", '8': "h",
	'7': "g", '6': "f", '5': "e", '4': "d", '3': "c", '2': "b"}

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

	high_map := make(map[string]hand_t)
	var highs []string
	one_map := make(map[string]hand_t)
	var ones []string
	two_map := make(map[string]hand_t)
	var twos []string
	three_map := make(map[string]hand_t)
	var threes []string
	full_map := make(map[string]hand_t)
	var fulls []string
	four_map := make(map[string]hand_t)
	var fours []string
	five_map := make(map[string]hand_t)
	var fives []string

	for _, hand := range input {
		tmp := strings.Split(hand, " ")
		var tmp_hand hand_t
		tmp_hand.cards = tmp[0]
		tmp_hand.bid, _ = strconv.Atoi(tmp[1])
		tmp_hand.type_h = parseType(tmp[0])
		tmp_hand.strength = cardsToStrenght(tmp_hand.cards)
		fmt.Println(tmp_hand)

		if tmp_hand.type_h == "five" {
			five_map[tmp_hand.strength] = tmp_hand
			fives = append(fives, tmp_hand.strength)
		} else if tmp_hand.type_h == "four" {
			four_map[tmp_hand.strength] = tmp_hand
			fours = append(fours, tmp_hand.strength)
		} else if tmp_hand.type_h == "full house" {
			full_map[tmp_hand.strength] = tmp_hand
			fulls = append(fulls, tmp_hand.strength)
		} else if tmp_hand.type_h == "three" {
			three_map[tmp_hand.strength] = tmp_hand
			threes = append(threes, tmp_hand.strength)
		} else if tmp_hand.type_h == "two pair" {
			two_map[tmp_hand.strength] = tmp_hand
			twos = append(twos, tmp_hand.strength)
		} else if tmp_hand.type_h == "one pair" {
			one_map[tmp_hand.strength] = tmp_hand
			ones = append(ones, tmp_hand.strength)
		} else { // high card
			high_map[tmp_hand.strength] = tmp_hand
			highs = append(highs, tmp_hand.strength)
		}
	}

	bid_sum := 0
	rank := 0
	evalBid(highs, high_map, &bid_sum, &rank)
	evalBid(ones, one_map, &bid_sum, &rank)
	evalBid(twos, two_map, &bid_sum, &rank)
	evalBid(threes, three_map, &bid_sum, &rank)
	evalBid(fulls, full_map, &bid_sum, &rank)
	evalBid(fours, four_map, &bid_sum, &rank)
	evalBid(fives, five_map, &bid_sum, &rank)
	fmt.Println(bid_sum)
}

func evalBid(strengths []string, str_map map[string]hand_t, bid_sum, rank *int) {
	sort.Strings(strengths)
	for _, hand := range strengths {
		*rank++
		*bid_sum += str_map[hand].bid * *rank
	}
}

func cardsToStrenght(hand string) (str string) {
	str = ""
	for _, card := range hand {
		str += strength_map[card]
	}
	return
}

func parseType(hand string) string {
	count := strings.Count(hand, string(hand[0]))
	if count == 5 {
		return "five"
	} else if count == 4 || count == 1 && strings.Count(hand, string(hand[1])) == 4 {
		return parseJ(hand, "four")
	} else if count == 3 {
		i := 1
		for ; hand[0] == hand[i]; i++ {
		}
		count = strings.Count(hand, string(hand[i]))
		if count == 2 {
			return parseJ(hand, "full house")
		} else {
			return parseJ(hand, "three")
		}
	} else if count == 2 {
		i := 1
		for ; hand[0] == hand[i]; i++ {
		}
		count = strings.Count(hand, string(hand[i]))
		if count == 3 {
			return parseJ(hand, "full house")
		} else if count == 2 {
			return parseJ(hand, "two pair")
		} else {
			i++
			for ; hand[0] == hand[i]; i++ {
			}
			count = strings.Count(hand, string(hand[i]))
			if count == 2 {
				return parseJ(hand, "two pair")
			} else {
				return parseJ(hand, "one pair")
			}
		}
	} else {
		count = strings.Count(hand, string(hand[1]))
		if count == 3 {
			return parseJ(hand, "three")
		} else if count == 2 {
			i := 2
			for ; hand[1] == hand[i]; i++ {
			}
			count = strings.Count(hand, string(hand[i]))
			if count == 2 {
				return parseJ(hand, "two pair")
			} else {
				return parseJ(hand, "one pair")
			}
		} else {
			count = strings.Count(hand, string(hand[2]))
			if count == 3 {
				return parseJ(hand, "three")
			} else if count == 2 || strings.Count(hand, string(hand[3])) == 2 {
				return parseJ(hand, "one pair")
			} else {
				return parseJ(hand, "high card")
			}
		}
	}

}

func parseJ(hand, type_h string) string {

	if count := strings.Count(hand, "J"); count == 5 || count == 4 {
		return "five"
	} else if count == 3 {
		return j3_map[type_h]
	} else if count == 2 {
		return j2_map[type_h]
	} else if count == 1 {
		return j1_map[type_h]
	}
	return type_h
}
