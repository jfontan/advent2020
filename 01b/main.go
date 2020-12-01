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

	for i := 0; i < len(numbers)-2; i++ {
		a := numbers[i]
		for j := i + 1; j < len(numbers)-1; j++ {
			b := numbers[j]
			for k := j + 1; k < len(numbers); k++ {
				c := numbers[k]
				if a+b+c == 2020 {
					fmt.Printf("%v * %v * %v= %v\n", a, b, c, a*b*c)
					return nil
				}
			}
		}
	}

	return fmt.Errorf("numbers not found")
}
