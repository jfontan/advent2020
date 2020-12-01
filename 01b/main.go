package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	err := solve()
	if err != nil {
		panic(err)
	}
}

func solve() error {
	f, err := os.Open("input")
	if err != nil {
		return err
	}
	defer f.Close()

	var numbers []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		n, err := strconv.Atoi(strings.TrimSpace(s.Text()))
		if err != nil {
			return err
		}
		numbers = append(numbers, n)
	}

	return find(numbers, 3, 2020)
}

func find(numbers []int, n, s int) error {
	counters := make([]int, n)

	for {
		sum := 0
		var vals []int
		for _, c := range counters {
			vals = append(vals, numbers[c])
			sum += numbers[c]
		}

		if sum == s {
			mul := 1
			for _, v := range vals {
				mul *= v
			}
			fmt.Printf("%v\n", mul)
			return nil
		}

		carry := true
		for i := range counters {
			if carry {
				counters[i]++
				carry = false
			}
			if counters[i] >= len(numbers) {
				counters[i] = 0
				carry = true
			}
		}

		if carry {
			break
		}
	}

	return fmt.Errorf("numbers not found")
}
