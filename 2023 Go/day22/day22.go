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
	} else { // go run . [inputfile]
		file = os.Args[1] + file
	}
	fmt.Println(file)
	file_in, _ := os.Open(file)
	defer file_in.Close()
	input := ParseInput(file_in)
	Part1_2(input)
}

func ParseInput(file_in *os.File) (input []string) {
	scanner := bufio.NewScanner(file_in)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

type cube_t struct {
	x, y, z [2]int
	heldBy  []int
	holds   []int
	chain   map[int]struct{}
}

func Part1_2(input []string) {
	m, s := Part1Parse(input)
	// lowering what can be lowered
	for z := 1; z < len(s); z++ {
		for y := range s[z] {
			for x := range s[z][y] {
				if s[z][y][x] != 0 {
					id := s[z][y][x]
					for CanLower(m[id], s) {
						m[id] = Lower(m[id], s)
					}
				}
			}
		}
	}
	// checking what holds(/is held) by what
	var tmp cube_t
	for k := range m {
		tmp = m[k]
		tmp.holds = Holds(tmp, s)
		m[k] = tmp
		//fmt.Printf("%v holds: %v\n", k, tmp.holds)
		for _, e := range m[k].holds {
			tmp = m[e]
			tmp.heldBy = append(m[e].heldBy, k)
			m[e] = tmp
			//fmt.Printf("  %v held by: %v\n", e, tmp.heldBy)
		}
	}
	// calculating sum for part1
	sum := 0
	for k := range m {
		if len(m[k].holds) > 0 {
			flag := false
			for _, e := range m[k].holds {
				if len(m[e].heldBy) <= 1 {
					flag = true
				}
			}
			if !flag {
				sum++
			}
		} else {
			sum++
		}
	}
	fmt.Println(sum)

	// Part2
	seen := make(map[int]struct{})
	for z := len(s) - 1; z > 0; z-- {
		for y := 0; y < len(s[z]); y++ {
			for x := 0; x < len(s[z][y]); x++ {
				if s[z][y][x] != 0 {
					if _, inseen := seen[s[z][y][x]]; !inseen {
						id := s[z][y][x]
						seen[id] = struct{}{}
						tmp = m[id]
						tmp.chain = make(map[int]struct{})
						if len(tmp.holds) > 0 {
							var idlist []int
							idlist = append(idlist, Destroy(tmp, s))
							var cur []cube_t
							cur = append(cur, tmp)
							var next []cube_t
							for cur != nil {
								for _, c := range cur {
									for _, h := range c.holds {
										if CanLower(m[h], s) {
											idlist = append(idlist, Destroy(m[h], s))
											tmp.chain[h] = struct{}{}
											next = append(next, m[h])
										}
									}
								}
								cur = next
								next = nil
							}
							Rebuild(m, s, idlist...)
						}
						m[id] = tmp
					}
				}
			}
		}
	}
	// calculating sum for part2
	sum2 := 0
	for _, v := range m {
		sum2 += len(v.chain)
	}
	fmt.Println(sum2)
}

func Part1Parse(input []string) (map[int]cube_t, [][][]int) {
	m := make(map[int]cube_t)
	id := 0
	maxX, maxY, maxZ := -1, -1, -1
	for _, line := range input {
		id++
		t := strings.Split(line, "~")
		l := strings.Split(t[0], ",")
		r := strings.Split(t[1], ",")
		var tx, ty, tz [2]int
		tx[0], _ = strconv.Atoi(l[0])
		tx[1], _ = strconv.Atoi(r[0])
		ty[0], _ = strconv.Atoi(l[1])
		ty[1], _ = strconv.Atoi(r[1])
		tz[0], _ = strconv.Atoi(l[2])
		tz[1], _ = strconv.Atoi(r[2])
		if tx[1] > maxX {
			maxX = tx[1]
		}
		if ty[1] > maxY {
			maxY = ty[1]
		}
		if tz[1] > maxZ {
			maxZ = tz[1]
		}
		tc := cube_t{}
		tc.x = tx
		tc.y = ty
		tc.z = tz
		m[id] = tc
	}
	maxX++
	maxY++
	maxZ++

	stack := make([][][]int, maxZ)
	for z := range stack {
		stack[z] = make([][]int, maxY)
		for y := range stack[z] {
			stack[z][y] = make([]int, maxX)
		}
	}
	for k, v := range m {
		for z := v.z[0]; z <= v.z[1]; z++ {
			for y := v.y[0]; y <= v.y[1]; y++ {
				for x := v.x[0]; x <= v.x[1]; x++ {
					stack[z][y][x] = k
				}
			}
		}
	}
	return m, stack
}

// Checks if cube can be lowered
func CanLower(c cube_t, s [][][]int) bool {
	id := s[c.z[0]][c.y[0]][c.x[0]]
	for x := c.x[0]; x <= c.x[1]; x++ {
		for y := c.y[0]; y <= c.y[1]; y++ {
			for z := c.z[0]; z <= c.z[1]; z++ {
				if s[z-1][y][x] != 0 && s[z-1][y][x] != id || z-1 == 0 {
					return false
				}
			}
		}
	}
	return true
}

// Lowers cube c, returning c with it's new position
func Lower(c cube_t, s [][][]int) cube_t {
	for x := c.x[0]; x <= c.x[1]; x++ {
		for y := c.y[0]; y <= c.y[1]; y++ {
			for z := c.z[0]; z <= c.z[1]; z++ {
				s[z-1][y][x] = s[z][y][x]
				if z == c.z[1] {
					s[z][y][x] = 0
				}
			}
		}
	}
	c.z[0]--
	c.z[1]--
	return c
}

// Returns an id list of what the cube holds
func Holds(c cube_t, s [][][]int) (idlist []int) {
	h := make(map[int]struct{})
	for x := c.x[0]; x <= c.x[1]; x++ {
		for y := c.y[0]; y <= c.y[1]; y++ {
			if s[c.z[1]+1][y][x] != 0 {
				h[s[c.z[1]+1][y][x]] = struct{}{}
			}
		}
	}
	for k := range h {
		idlist = append(idlist, k)
	}
	return
}

// Removes cube c from the stack s
func Destroy(c cube_t, s [][][]int) (id int) {
	id = s[c.z[0]][c.y[0]][c.x[0]]
	for x := c.x[0]; x <= c.x[1]; x++ {
		for y := c.y[0]; y <= c.y[1]; y++ {
			for z := c.z[0]; z <= c.z[1]; z++ {
				s[z][y][x] = 0
			}
		}
	}
	return
}

// Adds all ids in idlist back to the stack s
func Rebuild(m map[int]cube_t, s [][][]int, idlist ...int) {
	//id := s[c.z[0]][c.y[0]][c.x[0]]
	for i := 0; i < len(idlist); i++ {
		c := m[idlist[i]]
		for x := c.x[0]; x <= c.x[1]; x++ {
			for y := c.y[0]; y <= c.y[1]; y++ {
				for z := c.z[0]; z <= c.z[1]; z++ {
					s[z][y][x] = idlist[i]
				}
			}
		}
	}
}
