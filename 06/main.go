package main

import (
	"github.com/jfontan/advent2020/utils"
)

func main() {
	err := solve1()
	if err != nil {
		panic(err)
	}

	err = solve2()
	if err != nil {
		panic(err)
	}
}

func solve1() error {
	lines, err := utils.GetLines("06/input")
	if err != nil {
		return err
	}

	res := parse(lines)

	sum := 0
	for _, r := range res {
		sum += len(r)
	}

	println(sum)
	return nil
}

type responses map[rune]struct{}

func parse(lines []string) []responses {
	var all []responses
	r := make(responses)
	for _, l := range lines {
		if l == "" {
			all = append(all, r)
			r = make(responses)
			continue
		}

		for _, c := range l {
			r[c] = struct{}{}
		}
	}

	if len(r) != 0 {
		all = append(all, r)
	}

	return all
}

type group struct {
	number    int
	responses map[rune]int
}

func solve2() error {
	lines, err := utils.GetLines("06/input")
	if err != nil {
		return err
	}

	var groups []group
	g := group{responses: make(map[rune]int)}
	for _, l := range lines {
		if l == "" {
			groups = append(groups, g)
			g = group{responses: make(map[rune]int)}
			continue
		}

		for _, c := range l {
			g.responses[c]++
		}

		g.number++
	}

	groups = append(groups, g)

	sum := 0
	for _, g := range groups {
		for _, r := range g.responses {
			if r == g.number {
				sum++
			}
		}
	}

	println(sum)
	return nil
}
