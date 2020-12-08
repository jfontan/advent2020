package main

import (
	"bufio"
	"os"
	"strings"
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

var (
	Positions = map[rune]int{
		'F': 0,
		'B': 1,
		'L': 0,
		'R': 1,
	}
)

func solve1() error {
	lines, err := getLines("05/input")
	if err != nil {
		return err
	}

	var maxPos int
	for _, l := range lines {
		pos := seat(l)
		if pos > maxPos {
			maxPos = pos
		}
	}

	println(maxPos)

	return nil
}

func solve2() error {
	lines, err := getLines("05/input")
	if err != nil {
		return err
	}

	seats := make(map[int]struct{})
	for _, l := range lines {
		pos := seat(l)
		seats[pos] = struct{}{}
	}

	for i := 1; i < 127*8+6; i++ {
		_, ok := seats[i]
		if ok {
			continue
		}

		_, prev := seats[i-1]
		_, next := seats[i+1]

		if prev && next {
			println(i)
		}
	}

	return nil
}

func seat(l string) int {
	min, max := 0, 127
	for _, c := range l[:7] {
		pos := Positions[c]
		min, max = partition(min, max, pos)
	}
	row := min

	min, max = 0, 7
	for _, c := range l[7:] {
		pos := Positions[c]
		min, max = partition(min, max, pos)
	}
	column := min

	pos := row*8 + column
	return pos
}

func partition(min, max, pos int) (int, int) {
	if min == max {
		return min, min
	}

	size := max - min + 1
	half := size / 2

	if pos == 0 {
		return min, min + half - 1
	}

	return min + half, max
}

func getLines(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var lines []string
	for s.Scan() {
		lines = append(lines, strings.TrimSpace(s.Text()))
	}

	return lines, nil
}
