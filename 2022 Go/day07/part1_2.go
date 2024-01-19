package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type file_t struct {
	id   string
	size int
}

type dir_t struct {
	id   string
	par  *dir_t
	sub  []*dir_t
	file []file_t
}

func NewDir(id string, par *dir_t) (d dir_t) {
	d.id = id
	d.par = par
	return
}
func (d *dir_t) String() string {
	if d != nil {
		return d.id
	} else {
		return "<nil>"
	}
}

func Part1_2(input []string) {
	root := Parse1(input)
	//root.List(1)
	sum := 0
	cur := 70000000 - root.Sum(&sum)
	fmt.Println(sum) // result part1
	need := 30000000 - cur
	best := math.MaxInt32
	root.Free(need, &best)
	fmt.Println(best) // result part2
}

func Parse1(r []string) *dir_t {
	var dir, root *dir_t
	tdir := NewDir("/", nil)
	root = &tdir
	dir = root
	for i := 2; i < len(r); {
		if r[i][0] != '$' {
			for ; i < len(r) && r[i][0] != '$'; i++ {
				//fmt.Printf("line %v: [%v] @[%v]\n", i+1, r[i], dir)
				t1 := strings.Fields(r[i])
				if t1[0] == "dir" {
					d := NewDir(t1[1], dir)
					dir.sub = append(dir.sub, &d)
				} else {
					ti, _ := strconv.Atoi(t1[0])
					dir.file = append(dir.file, file_t{id: t1[1], size: ti})
				}
			}
		} else if /* r[i][0] == '$' && */ r[i][5] == '.' {
			//fmt.Printf("line %v: [%v] to [%v] @[%v]\n", i+1, r[i][2:], dir.par, dir)
			dir = dir.par
			i++
		} else { // if r[i] == "$ cd dirname"
			//fmt.Printf("line %v: [%v] from [%v] @[%v]\n", i+1, r[i][2:], dir, dir)
			t := strings.Fields(r[i])
			for _, s := range dir.sub {
				if s.id == t[2] {
					//fmt.Printf("$ ls\n")
					dir = s
					break
				}
			}
			i += 2 // skip "$ ls" line
		}
	}
	return root
}

func (d *dir_t) List(lv int) {
	t := ""
	for i := 0; i < lv; i++ {
		t += "-"
	}
	fmt.Printf("%s dir %v\n", t, d)
	for _, f := range d.file {
		fmt.Printf("-%s file %v\n", t, f.id)
	}
	for _, s := range d.sub {
		s.List(lv + 1)
	}
}

func (d *dir_t) Sum(sum *int) int {
	ssize := 0
	for _, s := range d.sub {
		ssize += s.Sum(sum)
	}
	fsize := 0
	for _, f := range d.file {
		fsize += f.size
	}
	if ssize+fsize <= 100000 {
		*sum += ssize + fsize
	}
	return ssize + fsize
}

func (d *dir_t) Free(need int, best *int) int {
	ssize := 0
	for _, s := range d.sub {
		ssize += s.Free(need, best)
	}
	fsize := 0
	for _, f := range d.file {
		fsize += f.size
	}
	if ssize+fsize >= need && ssize+fsize < *best {
		*best = ssize + fsize
	}
	return ssize + fsize
}
