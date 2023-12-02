package main

import (
	"advent2023/internal/aoc"
	"log"
	"strconv"
	"strings"
)

func solvePart1(input []string) (solution int) {
	sum := 0
	for _, line := range input {
		first := ""
		last := ""
		for _, c := range line {
			if _, err := strconv.Atoi(string(c)); err == nil {
				if first == "" {
					first = string(c)
				}
				last = string(c)
			}
		}
		value, err := strconv.Atoi(strings.Join([]string{first, last}, ""))
		if err != nil {
			log.Fatal(err)
		}
		sum = sum + value
	}
	return sum
}

var digits map[string]string = map[string]string{
	"one":   "1",
	"1":     "1",
	"two":   "2",
	"2":     "2",
	"three": "3",
	"3":     "3",
	"four":  "4",
	"4":     "4",
	"five":  "5",
	"5":     "5",
	"six":   "6",
	"6":     "6",
	"seven": "7",
	"7":     "7",
	"eight": "8",
	"8":     "8",
	"nine":  "9",
	"9":     "9",
}

type digitIndex struct {
	digit string
	index int
}

func solvePart2(input []string) (solution int) {
	sum := 0
	for _, line := range input {
		first := digitIndex{digit: "", index: 5000}
		last := digitIndex{digit: "", index: -1}
		for k := range digits {
			f := strings.Index(line, k)
			if f != -1 && f < first.index {
				first = digitIndex{digit: k, index: f}
			}
			l := strings.LastIndex(line, k)
			if l != -1 && l > last.index {
				last = digitIndex{digit: k, index: l}
			}
		}
		joinedDigits := strings.Join([]string{digits[first.digit], digits[last.digit]}, "")
		value, _ := strconv.Atoi(joinedDigits)
		sum = sum + value
	}
	return sum
}

func main() {
	day := 1
	aoc.SolveExample(day, aoc.Part1, solvePart1, 142)
	aoc.SolvePuzzle(day, aoc.Part1, solvePart1, 56397)
	aoc.SolveExample(day, aoc.Part2, solvePart2, 281)
	aoc.SolvePuzzle(day, aoc.Part2, solvePart2, 55701)
}
