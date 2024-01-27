package main

import (
	"container/heap"
	"image"
	"math"
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

const INFINITY int = math.MaxInt32

func Neighbors(cur, prev image.Point, bounds image.Rectangle, input [][]rune) (next []image.Point) {
	dir := [4]image.Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	if up := cur.Add(dir[0]); prev != up && up.In(bounds) && input[up.X][up.Y] <= input[cur.X][cur.Y]+1 {
		next = append(next, up)
	}
	if left := cur.Add(dir[1]); prev != left && left.In(bounds) && input[left.X][left.Y] <= input[cur.X][cur.Y]+1 {
		next = append(next, left)
	}
	if down := cur.Add(dir[2]); prev != down && down.In(bounds) && input[down.X][down.Y] <= input[cur.X][cur.Y]+1 {
		next = append(next, down)
	}
	if right := cur.Add(dir[3]); prev != right && right.In(bounds) && input[right.X][right.Y] <= input[cur.X][cur.Y]+1 {
		next = append(next, right)
	}
	return
}

func Dijkstra(start, end image.Point, input [][]rune) int {
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

	for len(openSet) > 0 {
		cur, cur_fs := openSet.GPop()

		if cur == end { // found end
			return cur_fs
		}
		delete(inQueue, cur)
		for _, n := range Neighbors(cur, cameFrom[cur], bounds, input) {
			if _, ok := gscore[n]; !ok {
				gscore[n] = INFINITY
			}
			tentative_gs := gscore[cur] + 1 // 1 == d(current, neighbor)
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
}
