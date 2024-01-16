package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"image"
	"math"
	"os"
	"strconv"
	"strings"
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

const (
	INFINITY      int = math.MaxInt32
	MINFINITY     int = math.MinInt32
	UNINITIALIZED int = -1
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
	Part2(input)
}

func ParseInput(filepath string) (input [][]rune) {
	file_in, _ := os.Open(filepath)
	defer file_in.Close()
	scanner := bufio.NewScanner(file_in)
	for scanner.Scan() {
		input = append(input, []rune(scanner.Text()))
	}
	return input
}

func Part2(input [][]rune) {

	bounds := image.Rect(0, 0, len(input), len(input[len(input)-1]))
	start := image.Point{0, 1}
	end := image.Point{len(input) - 1, len(input[len(input)-1]) - 2}

	prev := start
	last_i := start
	cur := image.Point{1, 1}

	adj := make(map[image.Point]map[image.Point]int)
	Compress(last_i, prev, cur, end, bounds, input, adj)

	debugprint := false
	if debugprint {
		for k := range adj {
			fmt.Println(k, adj[k])
		}
	} else {
		sum := 0
		for step1, dist := range adj[start] {
			sum += dist + Rec(dist, IJtoS(start.X, start.Y), step1, end, adj)
		}
		fmt.Println(sum)
		fmt.Println(total)
	}
}

func IJtoS(i, j int) string { return "(" + strconv.Itoa(i) + "," + strconv.Itoa(j) + ")" }

var total int = 0

func Rec(tot int, prev string, cur, end image.Point, adj map[image.Point]map[image.Point]int) int {
	if cur == end {
		if tot > total {
			total = tot
		}
		return 0
	}
	var val int
	hi_val := 0
	for neighbor, dist := range adj[cur] {
		n_str := IJtoS(neighbor.X, neighbor.Y)
		if strings.Contains(prev, n_str) {
			continue
		}
		val = dist + Rec(dist+tot, prev+n_str, neighbor, end, adj)

		if val > hi_val {
			hi_val = val
		}
	}
	return hi_val
}

func Compress(last_i, prev, cur, end image.Point, bounds image.Rectangle, input [][]rune, adj map[image.Point]map[image.Point]int) {
	for n := Neighbors(cur, prev, bounds, input); len(n) == 1; n = Neighbors(cur, prev, bounds, input) {
		prev = cur
		cur = n[0]
		if cur == end {
			break
		}
	}
	//fmt.Println("test")
	if _, ok := adj[last_i]; !ok {
		adj[last_i] = make(map[image.Point]int)
	}
	if adj[last_i][cur] == 0 {
		if _, ok := adj[cur]; !ok {
			adj[cur] = make(map[image.Point]int)
		}
		adj[last_i][cur] = Dij(last_i, cur, input)
		adj[cur][last_i] = adj[last_i][cur]
	} else {
		return
	}

	if cur == end {
		return
	}
	last_i = cur
	t := Neighbors(cur, prev, bounds, input)
	prev = cur
	for i := 0; i < len(t); i++ {
		if t[i] != cur {
			Compress(cur, prev, t[i], end, bounds, input, adj)
		}
	}
}

func Reconstruct_path(cur image.Point, cameFrom map[image.Point]image.Point, path map[image.Point]struct{}) int {
	cost := 0
	blank := image.Point{-1, -1}
	for cur := cameFrom[cur]; cur != blank; cur = cameFrom[cur] {
		path[cur] = struct{}{}
		cost++
	}
	return cost
}

func Neighbors(cur, prev image.Point, bounds image.Rectangle, input [][]rune) (next []image.Point) {
	dir := [4]image.Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	if up := cur.Add(dir[0]); prev != up && up.In(bounds) && input[up.X][up.Y] != '#' {
		next = append(next, up)
	}
	if left := cur.Add(dir[1]); prev != left && left.In(bounds) && input[left.X][left.Y] != '#' {
		next = append(next, left)
	}
	if down := cur.Add(dir[2]); prev != down && down.In(bounds) && input[down.X][down.Y] != '#' {
		next = append(next, down)
	}
	if right := cur.Add(dir[3]); prev != right && right.In(bounds) && input[right.X][right.Y] != '#' {
		next = append(next, right)
	}
	return
}

func Dij(start, end image.Point, input [][]rune) int {
	bounds := image.Rect(0, 0, len(input), len(input[len(input)-1]))

	openSet := PQ[image.Point]{}
	gscore := make(map[image.Point]int)           // starts set to INFINITY
	fscore := make(map[image.Point]int)           // starts set to INFINITY
	cameFrom := make(map[image.Point]image.Point) // starts empty
	inQueue := make(map[image.Point]struct{})     // starts empty

	gscore[start] = 0
	fscore[start] = 0 // + h(start)
	openSet.GPush(start, fscore[start])
	cameFrom[start] = image.Point{-1, -1}
	inQueue[start] = struct{}{}

	intersec := make(map[image.Point]struct{}) // starts empty

	for len(openSet) > 0 {
		cur, cur_fs := openSet.GPop()

		if cur == end { /*found end*/
			//fmt.Println(cur_fs)
			return cur_fs
		}
		delete(inQueue, cur)
		neigh := Neighbors(cur, cameFrom[cur], bounds, input)
		if len(neigh) >= 2 {
			intersec[cur] = struct{}{}
		}
		for _, n := range neigh {
			if _, ok := gscore[n]; !ok {
				gscore[n] = INFINITY
			}
			tentative_gs := gscore[cur] + 1 //d(current, neighbor)
			if tentative_gs < gscore[n] {
				cameFrom[n] = cur
				gscore[n] = tentative_gs
				fscore[n] = tentative_gs //+ h(neighbor)
				if _, in_open := inQueue[n]; !in_open {
					openSet.GPush(n, fscore[n])
					inQueue[n] = struct{}{}
				}
			}
		}
	}
	return -1
	/*
		path := make(map[image.Point]struct{})
		fmt.Println(Reconstruct_path(end, cameFrom, path))

		img := image.NewRGBA(bounds)
		for i := range input {
			for j := range input[i] {
				if (i == start.X && j == start.Y) || (i == end.X && j == end.Y) {
					img.Set(j, i, color.RGBA{000, 255, 000, 0xff})
				} else if _, ok := intersec[image.Point{i, j}]; ok {
					img.Set(j, i, color.RGBA{255, 000, 000, 0xff})
				} else if _, ok := path[image.Point{i, j}]; ok {
					img.Set(j, i, color.RGBA{000, 000, 255, 0xff})
				} else if _, ok := fscore[image.Point{i, j}]; ok {
					img.Set(j, i, color.White)
				} else {
					img.Set(j, i, color.Black)
				}
			}
		}

		f, _ := os.Create("image.gif")
		gif.Encode(f, img, nil)
		fmt.Println("image written")
	*/
}
