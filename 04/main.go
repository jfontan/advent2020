package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
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

const (
	TypeBYR = 1 << iota
	TypeIYR
	TypeEYR
	TypeHGT
	TypeHCL
	TypeECL
	TypePID
	TypeCID

	TypeValid = TypeBYR | TypeIYR | TypeEYR | TypeHGT | TypeHCL | TypeECL |
		TypePID
)

var (
	TypeMask = map[string]int{
		"byr": TypeBYR,
		"iyr": TypeIYR,
		"eyr": TypeEYR,
		"hgt": TypeHGT,
		"hcl": TypeHCL,
		"ecl": TypeECL,
		"pid": TypePID,
		"cid": TypeCID,
	}

	RegEYR = regexp.MustCompile(`^(\d+)(cm|in)$`)
	RegHCL = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	RegECL = regexp.MustCompile(`(amb|blu|brn|gry|grn|hzl|oth)$`)
	RegPID = regexp.MustCompile(`^[0-9]{9}$`)
)

type Data struct {
	Name  string
	Value string
}

func (d Data) Valid() bool {
	switch d.Name {
	case "byr":
		n, _ := strconv.Atoi(d.Value)
		if n < 1920 || n > 2002 {
			return false
		}
	case "iyr":
		n, _ := strconv.Atoi(d.Value)
		if n < 2010 || n > 2020 {
			return false
		}
	case "eyr":
		n, _ := strconv.Atoi(d.Value)
		if n < 2020 || n > 2030 {
			return false
		}
	case "hgt":
		m := RegEYR.FindStringSubmatch(d.Value)
		if m == nil {
			return false
		}
		n, _ := strconv.Atoi(m[1])
		if m[2] == "cm" {
			if n < 150 || n > 193 {
				return false
			}
		} else {
			if n < 50 || n > 76 {
				return false
			}
		}
	case "hcl":
		if !RegHCL.MatchString(d.Value) {
			return false
		}
	case "ecl":
		if !RegECL.MatchString(d.Value) {
			return false
		}
	case "pid":
		if !RegPID.MatchString(d.Value) {
			return false
		}
	}

	return true
}

type Passport []Data

func (p Passport) Mask() int {
	var mask int
	for _, d := range p {
		m, ok := TypeMask[d.Name]
		if ok {
			mask |= m
		}
	}
	return mask
}

func (p Passport) Valid() bool {
	for _, d := range p {
		if !d.Valid() {
			return false
		}
	}
	return true
}

func solve1() error {
	lines, err := getLines("04/input")
	if err != nil {
		return err
	}

	passports := parse(lines)

	var valid int
	for _, p := range passports {
		if p.Mask()&TypeValid == TypeValid {
			valid++
		}
	}

	println(valid)

	return nil
}

func solve2() error {
	lines, err := getLines("04/input")
	if err != nil {
		return err
	}

	passports := parse(lines)

	var valid int
	for _, p := range passports {
		if p.Mask()&TypeValid != TypeValid {
			continue
		}

		if p.Valid() {
			valid++
		}
	}

	println(valid)

	return nil
}

func parse(lines []string) []Passport {
	var passports []Passport

	var passport Passport
	for _, l := range lines {
		if l == "" {
			if len(passport) > 0 {
				passports = append(passports, passport)
			}
			passport = nil
			continue
		}

		parts := strings.Split(l, " ")
		for _, part := range parts {
			p := strings.Split(part, ":")
			if len(p) != 2 {
				continue
			}

			d := Data{
				Name:  p[0],
				Value: p[1],
			}
			passport = append(passport, d)
		}
	}

	if len(passport) > 0 {
		passports = append(passports, passport)
	}

	return passports
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
