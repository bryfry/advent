package main

import (
	"advent/internal/aoc"
	"log"
	"strconv"
	"strings"
)

type direction int

const (
	unknown    direction = iota
	increasing direction = iota
	decreasing direction = iota
	volitile   direction = iota
)

type report struct {
	levels     []int
	direction  direction
	safe       bool
	violations int
}

func parse(line string) report {
	sl := strings.Split(line, " ")

	r := report{
		levels:     make([]int, len(sl), len(sl)),
		direction:  unknown,
		safe:       true,
		violations: 0,
	}

	previousLevel := -1
	for i, s := range sl {

		level, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("failed to parse level=%q: %s\n", s, err)
		}
		r.levels[i] = level

		if previousLevel == -1 {
			previousLevel = level
			continue
		}

		direction := unknown
		if level > previousLevel {
			direction = increasing
		}
		if level < previousLevel {
			direction = decreasing
		}
		if r.direction == unknown {
			r.direction = direction
		}
		if (r.direction == increasing && direction == decreasing) ||
			(r.direction == decreasing && direction == increasing) {
			r.direction = volitile
			r.safe = false
			r.violations += 1
		}

		diff := level - previousLevel
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			r.safe = false
			r.violations += 1
		}

		previousLevel = level
	}
	//fmt.Printf("%+v\n", r)
	return r
}

func solvePart1(input []string) (solution int) {
	safeReportSum := 0
	for _, l := range input {
		r := parse(l)
		if r.safe {
			safeReportSum += 1
		}
	}
	return safeReportSum
}

func solvePart2(input []string) (solution int) {
	safeReportSum := 0
	for _, l := range input {
		r := parse(l)
		if r.safe || r.violations == 1 {
			safeReportSum += 1
		}
	}
	return safeReportSum
}

func main() {
	day := aoc.Day2
	aoc.SolveExample(day, aoc.Part1, solvePart1, 2)
	aoc.SolvePuzzle(day, aoc.Part1, solvePart1, 564)
	aoc.SolveExample(day, aoc.Part2, solvePart2, 4)
	aoc.SolvePuzzle(day, aoc.Part2, solvePart2, aoc.UnknownExpected)
}
