package utils

import (
	"bufio"
	"os"
	"strings"
)

func GetLines(file string) ([]string, error) {
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
