package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type pqi[T any] struct {
	v T
	p int
}

type PQ[T any] []pqi[T]

func (q PQ[_]) Len() int           { return len(q) }
func (q PQ[_]) Less(i, j int) bool { return q[i].p < q[j].p }
func (q PQ[_]) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *PQ[T]) Push(x any)        { *q = append(*q, x.(pqi[T])) }
func (q *PQ[_]) Pop() (x any)      { x, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]; return x }
func (q *PQ[T]) GPush(v T, p int)  { heap.Push(q, pqi[T]{v, p}) }
func (q *PQ[T]) GPop() (T, int)    { x := heap.Pop(q).(pqi[T]); return x.v, x.p }

type Node struct {
	i, j int
}

const INFINITY int = math.MinInt32

func main() {
	file := ".txt"
	if len(os.Args) == 1 {
		file = "sample" + file
	} else { // go run . [inputfile]
		file = os.Args[1] + file
	}
	fmt.Println(file)
	file_in, _ := os.Open(file)
	defer file_in.Close()
	input := ParseInput(file_in)
	Part1(input)
}

func ParseInput(file_in *os.File) (input [][]rune) {
	scanner := bufio.NewScanner(file_in)
	for scanner.Scan() {
		input = append(input, []rune(scanner.Text()))
	}
	return input
}

func Part1(input [][]rune) {
	queue := PQ[Node]{}
	start := Node{0, 1}
	end := Node{len(input) - 1, len(input[len(input)-1]) - 2}
	queue.GPush(start, 0)

	prev := make(map[Node]Node)
	next := make(map[Node][]Node)
	empty := Node{}
	prev[start] = empty

	path := make(map[Node](map[Node]struct{}))

	end_f := 0
	for len(queue) > 0 {
		node, fscore := queue.GPop()

		if node == end && fscore > end_f {
			end_f = fscore
		}

		if _, ok := path[node]; !ok {
			path[node] = make(map[Node]struct{})
		}
		if prev[node] != empty {
			if _, ok := path[node][prev[node]]; !ok {
				for k := range path[prev[node]] {
					path[node][k] = struct{}{}
				}
				path[node][prev[node]] = struct{}{}
			}
		}
		for _, n := range Neighbors(node, input, prev) {
			prev[n] = node
			if _, ok := path[node][n]; !ok {
				next[node] = append(next[node], n)
				queue.GPush(n, fscore+1)
			}
		}
	}
	fmt.Println(end_f)
}

func Neighbors(n Node, input [][]rune, prev map[Node]Node) (next []Node) {
	i, j := n.i, n.j
	u, d, l, r := i-1, i+1, j-1, j+1
	if u > 0 && input[u][j] != '#' && input[u][j] != 'v' {
		node_u := Node{u, j}
		if prev[n] != node_u {
			next = append(next, node_u)
		}
	}
	if d < len(input) && input[d][j] != '#' {
		node_d := Node{d, j}
		if prev[n] != node_d {
			next = append(next, node_d)
		}
	}
	if l > 0 && input[i][l] != '#' && input[i][l] != '>' {
		node_l := Node{i, l}
		if prev[n] != node_l {
			next = append(next, node_l)
		}
	}
	if r < len(input[i]) && input[i][r] != '#' {
		node_r := Node{i, r}
		if prev[n] != node_r {
			next = append(next, node_r)
		}
	}
	return
}
