package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var passwordReg = regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w+)$`)

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
	f, err := os.Open("02/input")
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var valid int
	for s.Scan() {
		line := s.Text()
		min, max, letter, pass, err := parse(line)
		if err != nil {
			return err
		}

		var count int
		for _, l := range pass {
			if l == rune(letter) {
				count++
			}
		}

		if count >= min && count <= max {
			valid++
		}
	}

	println(valid)

	return nil
}

func parse(s string) (int, int, byte, string, error) {
	m := passwordReg.FindStringSubmatch(s)
	if m == nil {
		return 0, 0, ' ', "", fmt.Errorf("malformed line: %s", s)
	}

	min, err := strconv.Atoi(m[1])
	if err != nil {
		return 0, 0, ' ', "", err
	}

	max, err := strconv.Atoi(m[2])
	if err != nil {
		return 0, 0, ' ', "", err
	}

	letter := m[3][0]
	pass := m[4]

	return min, max, letter, pass, nil
}

func solve2() error {
	f, err := os.Open("02/input")
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var valid int
	for s.Scan() {
		line := s.Text()
		min, max, letter, pass, err := parse(line)
		if err != nil {
			return err
		}

		if min-1 > len(pass) || max-1 > len(pass) {
			continue
		}

		var count int
		if pass[min-1] == letter {
			count++
		}
		if pass[max-1] == letter {
			count++
		}

		if count == 1 {
			valid++
		}
	}

	println(valid)

	return nil
}
