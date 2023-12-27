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
	//Part1(input)
	Part2(input)
}

func ParseInput(file_in *os.File) (input []string) {
	scanner := bufio.NewScanner(file_in)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func Part1Parse(input []string) (map[string]string, []xmas_t) {
	var workflow []string
	var rating []string
	isWKF := true
	for _, line := range input {
		if line == "" {
			isWKF = false
			continue
		}
		if isWKF {
			workflow = append(workflow, line)
		} else {
			rating = append(rating, line)
		}
	}

	wfw_m := make(map[string]string)
	for _, w := range workflow {
		tmp := strings.Split(w[:len(w)-1], "{")
		wfw_m[tmp[0]] = tmp[1]
	}

	var xmas []xmas_t
	for i, r := range rating {
		xmas = append(xmas, xmas_t{})
		field := strings.Split(r[1:len(r)-1], ",")
		for _, f := range field {
			switch f[0] {
			case 'x':
				xmas[i].x, _ = strconv.Atoi(f[2:])
			case 'm':
				xmas[i].m, _ = strconv.Atoi(f[2:])
			case 'a':
				xmas[i].a, _ = strconv.Atoi(f[2:])
			case 's':
				xmas[i].s, _ = strconv.Atoi(f[2:])
			}
		}
	}
	return wfw_m, xmas
}

type xmas_t struct {
	x, m, a, s int
}

func WkfParse1(w string, xmas xmas_t) string {
	rules := strings.Split(w, ",")
	for _, r := range rules {
		if strings.Contains(r, ":") {
			split_r := strings.Split(r, ":")
			switch split_r[0][0] {
			case 'x':
				if split_r[0][1] == '<' {
					ti, _ := strconv.Atoi(split_r[0][2:])
					if xmas.x < ti {
						return split_r[1]
					}
				} else { // split_r[0][1] == '>'
					ti, _ := strconv.Atoi(split_r[0][2:])
					if xmas.x > ti {
						return split_r[1]
					}
				}
			case 'm':
				if split_r[0][1] == '<' {
					ti, _ := strconv.Atoi(split_r[0][2:])
					if xmas.m < ti {
						return split_r[1]
					}
				} else { // split_r[0][1] == '>'
					ti, _ := strconv.Atoi(split_r[0][2:])
					if xmas.m > ti {
						return split_r[1]
					}
				}
			case 'a':
				if split_r[0][1] == '<' {
					ti, _ := strconv.Atoi(split_r[0][2:])
					if xmas.a < ti {
						return split_r[1]
					}
				} else { // split_r[0][1] == '>'
					ti, _ := strconv.Atoi(split_r[0][2:])
					if xmas.a > ti {
						return split_r[1]
					}
				}
			case 's':
				if split_r[0][1] == '<' {
					ti, _ := strconv.Atoi(split_r[0][2:])
					if xmas.s < ti {
						return split_r[1]
					}
				} else { // split_r[0][1] == '>'
					ti, _ := strconv.Atoi(split_r[0][2:])
					if xmas.s > ti {
						return split_r[1]
					}
				}
			}
		} else {
			return r
		}
	}
	return ""
}

func Part1(input []string) {
	wfw_m, xmas := Part1Parse(input)
	sum := 0
	for i := range xmas {
		cur := "in"
		for cur != "R" && cur != "A" {
			cur = WkfParse1(wfw_m[cur], xmas[i])
			if cur == "A" {
				sum += xmas[i].x + xmas[i].m + xmas[i].a + xmas[i].s
			}
		}
	}
	fmt.Println(sum)
}

type xmas2_t struct {
	x, m, a, s [2]int
	inst       string
}

func (x xmas2_t) Init(s string, c ...xmas2_t) xmas2_t {
	if len(c) == 0 {
		x.x[0] = 1
		x.x[1] = 4000
		x.m[0] = 1
		x.m[1] = 4000
		x.a[0] = 1
		x.a[1] = 4000
		x.s[0] = 1
		x.s[1] = 4000
		x.inst = s
		return x
	} else {
		c[0].inst = s
		return c[0]
	}
}

func (r xmas2_t) String() string {
	if r.x[0] != 0 {
		return fmt.Sprintf("x(%v) m(%v) a(%v) s(%v) ➜  %v\n", r.x, r.m, r.a, r.s, r.inst)
	} else {
		return fmt.Sprintf("➜  %v\n", r.inst)
	}
}

func (c xmas2_t) GetRating() (r uint64) {
	x := uint64(c.x[1] - c.x[0] + 1)
	m := uint64(c.m[1] - c.m[0] + 1)
	a := uint64(c.a[1] - c.a[0] + 1)
	s := uint64(c.s[1] - c.s[0] + 1)

	r = x * m * a * s
	return
}

func Part2(input []string) {
	wfw, _ := Part1Parse(input)

	var cur []xmas2_t
	var next []xmas2_t
	var app []xmas2_t
	cur = append(cur, xmas2_t{}.Init("in"))
	for cur != nil {
		for _, e := range cur {
			tn, ta := WkfParse2(wfw[e.inst], e)
			next = append(next, tn...)
			app = append(app, ta...)
		}
		cur, next = next, cur
		next = nil
	}

	sum := uint64(0)
	for _, a := range app {
		sum += a.GetRating()
	}
	fmt.Println(sum)
}

func WkfParse2(w string, c xmas2_t) (ruleset []xmas2_t, app []xmas2_t) {
	//fmt.Println(c.inst, w)
	rules := strings.Split(w, ",")
	for _, r := range rules {
		s1, e1 := 1, 4000
		s2, e2 := 1, 4000
		if strings.Contains(r, ":") {
			split_r := strings.Split(r, ":")
			rule := split_r[0]
			result := split_r[1]
			val, _ := strconv.Atoi(rule[2:])
			if rule[1] == '>' {
				s1 = val + 1
				e2 = val
			} else { // rule[1] == '<'
				e1 = val - 1
				s2 = val
			}
			switch rule[0] {
			case 'x':
				if c.x[0] > s1 {
					s1 = c.x[0]
				}
				if c.x[1] < e1 {
					e1 = c.x[1]
				}
				if c.x[0] > s2 {
					s2 = c.x[0]
				}
				if c.x[1] < e2 {
					e2 = c.x[1]
				}
				if result != "A" && result != "R" {
					ruleset = append(ruleset, xmas2_t{[2]int{s1, e1}, c.m, c.a, c.s, result})
				} else if result == "A" {
					app = append(app, xmas2_t{[2]int{s1, e1}, c.m, c.a, c.s, result})
				}
				c.x[0] = s2
				c.x[1] = e2
			case 'm':
				if c.m[0] > s1 {
					s1 = c.m[0]
				}
				if c.m[1] < e1 {
					e1 = c.m[1]
				}
				if c.m[0] > s2 {
					s2 = c.m[0]
				}
				if c.m[1] < e2 {
					e2 = c.m[1]
				}
				if result != "A" && result != "R" {
					ruleset = append(ruleset, xmas2_t{c.x, [2]int{s1, e1}, c.a, c.s, result})
				} else if result == "A" {
					app = append(app, xmas2_t{c.x, [2]int{s1, e1}, c.a, c.s, result})
				}
				c.m[0] = s2
				c.m[1] = e2
			case 'a':
				if c.a[0] > s1 {
					s1 = c.a[0]
				}
				if c.a[1] < e1 {
					e1 = c.a[1]
				}
				if c.a[0] > s2 {
					s2 = c.a[0]
				}
				if c.a[1] < e2 {
					e2 = c.a[1]
				}
				if result != "A" && result != "R" {
					ruleset = append(ruleset, xmas2_t{c.x, c.m, [2]int{s1, e1}, c.s, result})
				} else if result == "A" {
					app = append(app, xmas2_t{c.x, c.m, [2]int{s1, e1}, c.s, result})
				}
				c.a[0] = s2
				c.a[1] = e2
			case 's':
				if c.s[0] > s1 {
					s1 = c.s[0]
				}
				if c.s[1] < e1 {
					e1 = c.s[1]
				}
				if c.s[0] > s2 {
					s2 = c.s[0]
				}
				if c.s[1] < e2 {
					e2 = c.s[1]
				}
				if result != "A" && result != "R" {
					ruleset = append(ruleset, xmas2_t{c.x, c.m, c.a, [2]int{s1, e1}, result})

				} else if result == "A" {
					app = append(app, xmas2_t{c.x, c.m, c.a, [2]int{s1, e1}, result})
				}
				c.s[0] = s2
				c.s[1] = e2
			}
		} else {
			if r != "A" && r != "R" {
				ruleset = append(ruleset, xmas2_t{}.Init(r, c))
			} else if r == "A" {
				app = append(app, xmas2_t{}.Init("A", c))
			}
		}
	}
	//fmt.Printf("r:\n%v a:\n%v\n", ruleset, app)
	return
}
