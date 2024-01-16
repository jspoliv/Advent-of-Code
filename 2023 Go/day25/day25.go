package main

import (
	"bufio"
	"fmt"
	"os"
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
	Part1(input)
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

func Part1Parse(input []string) (map[string][]string, map[string]struct{}) {
	graph := make(map[string][]string)
	pair := make(map[string]struct{})

	for _, line := range input {
		f := strings.Fields(line[:3] + line[4:])
		for i := 1; i < len(f); i++ {
			graph[f[0]] = append(graph[f[0]], f[i])
			graph[f[i]] = append(graph[f[i]], f[0])
			if _, ok := pair[f[i]+" "+f[0]]; !ok {
				pair[f[0]+" "+f[i]] = struct{}{}
			}
		}
	}
	return graph, pair
}

func Part1(input []string) {
	_, p := Part1Parse(input)
	fmt.Println(len(p))

}

func In(s string, l []string) bool {
	for _, e := range l {
		if s == e {
			return true
		}
	}
	return false
}

/*
p := [2]string{"hfx", "pzl"}
	DelBetween(p, g)
	p = [2]string{"bvb", "cmg"}
	DelBetween(p, g)
	p = [2]string{"nvd", "jqt"}
	DelBetween(p, g)
	PrintG(g)
*/

func DelBetween(pair []string, graph map[string][]string) {
	//fmt.Printf("%v: %v\n", pair[0], graph[pair[0]])
	for i, l := range graph[pair[0]] {
		if l == pair[1] {
			graph[pair[0]] = append(graph[pair[0]][:i], graph[pair[0]][i+1:]...)
		}
	}
	//fmt.Printf("%v: %v\n", pair[0], graph[pair[0]])
	//fmt.Printf("%v: %v\n", pair[1], graph[pair[1]])
	for i, l := range graph[pair[1]] {
		if l == pair[0] {
			graph[pair[1]] = append(graph[pair[1]][:i], graph[pair[1]][i+1:]...)
		}
	}
	//fmt.Printf("%v: %v\n", pair[1], graph[pair[1]])
}

func PrintG(graph map[string][]string) {
	for k, v := range graph {
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Println()
}
