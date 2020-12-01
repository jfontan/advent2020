package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	err := find()
	if err != nil {
		panic(err)
	}
}

func find() error {
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

	for i := 0; i < len(numbers)-1; i++ {
		a := numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			b := numbers[j]
			if a+b == 2020 {
				fmt.Printf("%v * %v = %v\n", a, b, a*b)
				return nil
			}
		}
	}

	return fmt.Errorf("numbers not found")
}
