package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

var LCM uint64 = 1

func Part2(input []string) {
	g := Parse1(input)
	for _, m := range g {
		LCM *= uint64(m.test)
	}
	counter := make(map[string]int)
	for m := range g {
		counter[m] = 0
	}
	del := make(map[string][]uint64)
	for round := 0; round < 10000; round++ {
		for i := 0; i < len(g); i++ {
			for _, item := range g[strconv.Itoa(i)].items {
				val, passTo := g[strconv.Itoa(i)].Inspect2(item)
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

func (m Monkey_t) Inspect2(item uint64) (uint64, string) {
	val, err := m.operate2(item)
	if err != nil {
		log.Fatal(err)
	}
	if val%uint64(m.test) == 0 {
		return val, m.passT
	} else {
		return val, m.passF
	}
}

func (m Monkey_t) operate2(item uint64) (val uint64, err error) {
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
		return (A * B) % LCM, nil // lower worry
	} else { // m.opS == "+"
		if A+B < A || A+B < B {
			return 0, errors.New("A+B overflow")
		}
		return (A + B) % LCM, nil // lower worry
	}
}
