package main

import (
	"advent/internal/aoc"
	"regexp"
	"strconv"
	"strings"
)

const (
	dontRegex = `don't\(\).*?(do\(\)|$)`
	mulRegex  = `mul\([0-9]+,[0-9]+\)`
	mulPrefix = "mul("
	mulSuffix = ")"
	mulSep    = ","
	findAll   = -1
)

func mul(s string) int {
	s = strings.TrimPrefix(s, mulPrefix)
	s = strings.TrimSuffix(s, mulSuffix)
	numbers := strings.Split(s, mulSep)
	// TODO assert length, check errors
	l, _ := strconv.Atoi(numbers[0])
	r, _ := strconv.Atoi(numbers[1])
	return l * r
}

func solvePart1(input []string) (solution int) {
	r := regexp.MustCompile(mulRegex)

	sum := 0
	for _, line := range input {
		matches := r.FindAllString(line, findAll)
		for _, m := range matches {
			v := mul(m)
			sum += v
			//fmt.Printf("%s=%d\n", m, v)
		}
	}
	return sum
}

func solvePart2(input []string) (solution int) {
	d := regexp.MustCompile(dontRegex)
	r := regexp.MustCompile(mulRegex)

	sum := 0
	for _, line := range input {
		line = d.ReplaceAllString(line, "")
		matches := r.FindAllString(line, findAll)
		for _, m := range matches {
			v := mul(m)
			sum += v
		}
	}
	return sum
}

func main() {
	day := aoc.Day3
	aoc.SolveExample(day, aoc.Part1, solvePart1, 161)
	aoc.SolvePuzzle(day, aoc.Part1, solvePart1, 189600467)
	aoc.SolveExample(day, aoc.Part2, solvePart2, 48)
	aoc.SolvePuzzle(day, aoc.Part2, solvePart2, 107069718)
}
