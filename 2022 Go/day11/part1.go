package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Monkey_t struct {
	id            string
	items         []uint64
	opA, opS, opB string
	test          int
	passT, passF  string
}

func FindnAdd(id string, item uint64, g map[string]Monkey_t) bool {
	for e := range g {
		if e == id {
			tmp := g[id]
			tmp.items = append(tmp.items, item)
			g[id] = tmp
			return true
		}
	}
	return false
}

func (m *Monkey_t) Remove(item uint64) bool {
	for i := 0; i < len(m.items); i++ {
		if m.items[i] == item {
			m.items = append(m.items[:i], m.items[i+1:]...)
			return true
		}
	}
	return false
}

func (m Monkey_t) Inspect(item uint64) (uint64, string) {
	val, _ := m.operate(item)
	if val%uint64(m.test) == 0 {
		return val, m.passT
	} else {
		return val, m.passF
	}
}

func (m Monkey_t) operate(item uint64) (val uint64, err error) {
	var A, B uint64
	if m.opA == "old" {
		A = item
	} else {
		A, _ = strconv.ParseUint(m.opA, 10, 64)
	}
	if m.opB == "old" {
		B = item
	} else {
		B, _ = strconv.ParseUint(m.opB, 10, 64)
	}
	if m.opS == "*" {
		if A*B < A || A*B < B {
			return 0, errors.New("A*B overflow")
		}
		val = uint64(math.Floor((float64(A*B) / float64(3)))) // lower worry
		return val, nil
	} else { // m.opS == "+"
		if A+B < A || A+B < B {
			return 0, errors.New("A+B overflow")
		}
		val = uint64(math.Floor((float64(A+B) / float64(3)))) // lower worry
		return val, nil
	}
}

func (m Monkey_t) String() string {
	return fmt.Sprintf("id: %v \nitems: %v \nnew = %v %v %v \ntest: %v ? true: %v : false: %v\n",
		m.id, m.items, m.opA, m.opS, m.opB, m.test, m.passT, m.passF)
}

func Parse1(input []string) map[string]Monkey_t {
	var mky Monkey_t
	group := make(map[string]Monkey_t)
	var t1 string
	var t2 []string
	var ti uint64
	for _, row := range input {
		if strings.Contains(row, "Monkey") {
			mky = Monkey_t{}
			t1, _ = strings.CutSuffix(row, ":")
			t2 = strings.Fields(t1)
			mky.id = t2[len(t2)-1]
		} else if strings.Contains(row, "items") {
			t1, _ = strings.CutPrefix(row, "  Starting items: ")
			t2 = strings.Split(t1, ", ")
			for _, t := range t2 {
				ti, _ = strconv.ParseUint(t, 10, 64)
				mky.items = append(mky.items, ti)
			}
		} else if strings.Contains(row, "Operation") {
			t2 = strings.Fields(row)
			mky.opB = t2[len(t2)-1]
			mky.opS = t2[len(t2)-2]
			mky.opA = t2[len(t2)-3]
		} else if strings.Contains(row, "divisible") {
			t2 = strings.Fields(row)
			mky.test, _ = strconv.Atoi(t2[len(t2)-1])
		} else if strings.Contains(row, "true") {
			t2 = strings.Fields(row)
			mky.passT = t2[len(t2)-1]
		} else if strings.Contains(row, "false") {
			t2 = strings.Fields(row)
			mky.passF = t2[len(t2)-1]
			group[mky.id] = mky
		}
	}
	return group
}

func Part1(input []string) {
	g := Parse1(input)
	counter := make(map[string]int)
	for m := range g {
		counter[m] = 0
	}
	del := make(map[string][]uint64)
	for round := 0; round < 20; round++ {
		for i := 0; i < len(g); i++ {
			for _, item := range g[strconv.Itoa(i)].items {
				val, passTo := g[strconv.Itoa(i)].Inspect(item)
				//fmt.Printf("%v item: %v >> %v\n", strconv.Itoa(i), item, val)
				counter[strconv.Itoa(i)]++
				FindnAdd(passTo, val, g)
				tmp := del[strconv.Itoa(i)]
				tmp = append(tmp, item)
				del[strconv.Itoa(i)] = tmp
			}
			for id, item := range del {
				for _, e := range item {
					mky := g[id]
					mky.Remove(e)
					g[id] = mky
				}
			}
			for id := range del {
				delete(del, id)
			}
		}
	}
	//for i := 0; i < len(g); i++ {
	//	fmt.Println(g[strconv.Itoa(i)].items)
	//}
	h1, h2 := -1, -2
	for i := 0; i < len(counter); i++ {
		if counter[strconv.Itoa(i)] > h1 {
			h2 = h1
			h1 = counter[strconv.Itoa(i)]
		} else if counter[strconv.Itoa(i)] > h2 {
			h2 = counter[strconv.Itoa(i)]
		}
	}
	fmt.Println(h1 * h2)
}
