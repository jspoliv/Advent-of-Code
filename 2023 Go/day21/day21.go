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
	} else {
		file = os.Args[1] + file
	}
	fmt.Println(file)
	file_in, _ := os.Open(file)
	defer file_in.Close()
	input := ParseInput(file_in)
	Part1(input)
	//Part2(input)
}

func ParseInput(file_in *os.File) (input [][]rune) {
	scanner := bufio.NewScanner(file_in)
	for scanner.Scan() {
		input = append(input, []rune(scanner.Text()))
	}
	return input
}

func Part1(input [][]rune) {
	m, g, si, sj := Part1Parse(input)
	dist, _ := Dijkstra(g, g.ids[IJtoS(si, sj)])

	sum := 0
	for k := range m {
		if d := dist[g.ids[k]]; /*d <= 10 &&*/ d%2 == 1 {
			sum++
		}
	}
	fmt.Println(sum)

	w := make([][]int, len(input))
	for x := range input {
		w[x] = make([]int, len(input[x]))
		for y := range input[x] {
			w[x][y] = -1
		}
	}
	for k := range m {
		//fmt.Printf("| %v %v ", k, dist[g.ids[k]])
		ti, tj := StoIJ(k)
		//fmt.Printf("%v ", string(input[ti][tj]))
		w[ti][tj] = dist[g.ids[k]]
	}
	//fmt.Println()
	for _, line := range w {
		for _, cell := range line {
			fmt.Printf("%2.d ", cell)
		}
		fmt.Println()
	}

}

func Part1Parse(input [][]rune) (m map[string]Vertex, g sg, si, sj int) {
	m = make(map[string]Vertex)
	count := Vertex(0)
	for i := range input {
		for j := range input[i] {
			if input[i][j] != '#' {
				cur := IJtoS(i, j)
				m[cur] = count
				count++
				if input[i][j] == 'S' {
					si = i
					sj = j
				}
			}
		}
	}

	g = newsg(m)

	for k := range m {
		i, j := StoIJ(k)
		up, down, left, right := i-1, i+1, j-1, j+1
		if up > 0 && input[up][j] != '#' {
			g.edge(k, IJtoS(up, j), 1)
			g.edge(IJtoS(up, j), k, 1)
		}
		if left > 0 && input[i][left] != '#' {
			g.edge(k, IJtoS(i, left), 1)
			g.edge(IJtoS(i, left), k, 1)
		}
		if down < len(input) && input[down][j] != '#' {
			g.edge(k, IJtoS(down, j), 1)
			g.edge(IJtoS(down, j), k, 1)
		}
		if right < len(input[i]) && input[i][right] != '#' {
			g.edge(k, IJtoS(i, right), 1)
			g.edge(IJtoS(i, right), k, 1)
		}
	}

	return m, g, si, sj
}

func IJtoS(i, j int) string { return strconv.Itoa(i) + "," + strconv.Itoa(j) }
func StoIJ(s string) (i, j int) {
	t := strings.Split(s, ",")
	i, _ = strconv.Atoi(t[0])
	j, _ = strconv.Atoi(t[1])
	return
}
