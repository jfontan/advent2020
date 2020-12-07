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

func solve1() error {
	lines, err := getLines("03/input")
	if err != nil {
		return err
	}

	trees := countTrees(lines, 3)

	println(trees)
	return nil
}

func solve2() error {
	lines, err := getLines("03/input")
	if err != nil {
		return err
	}

	trees := countTrees(lines, 1)
	trees *= countTrees(lines, 3)
	trees *= countTrees(lines, 5)
	trees *= countTrees(lines, 7)

	var t int
	for i := 0; i < len(lines); i += 2 {
		if lines[i][(i/2)%len(lines[i])] == '#' {
			t++
		}
	}

	trees *= t

	println(trees)
	return nil
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

func countTrees(lines []string, slope int) int {
	var trees int
	for i, l := range lines {
		pos := (i * slope) % len(l)
		if l[pos] == '#' {
			trees++
		}
	}
	return trees
}
