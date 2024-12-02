package main

import (
	"advent/internal/aoc"
	"log"
	"sort"
	"strconv"
	"strings"
)

func parse(line string) (left, right int) {
	sl := strings.Split(line, "   ")
	if len(sl) != 2 {
		log.Fatal("failed to parse line")
	}

	l, err := strconv.Atoi(sl[0])
	if err != nil {
		log.Fatalf("failed to parse left=%q: %s\n", sl[0], err)
	}

	r, err := strconv.Atoi(sl[1])
	if err != nil {
		log.Fatalf("failed to parse left=%q: %s\n", sl[1], err)
	}
	return l, r
}

func solvePart1(input []string) (solution int) {
	left, right := make([]int, len(input)), make([]int, len(input))
	for i, line := range input {
		left[i], right[i] = parse(line)
	}

	sort.Ints(left)
	sort.Ints(right)

	dist := 0
	for i := range left {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff
		}
		dist += diff
	}
	return dist
}

func solvePart2(input []string) (solution int) {

	left := make([]int, len(input))
	rightCount := make(map[int]int, len(input))

	for i, line := range input {
		l, r := parse(line)
		left[i] = l
		v, ok := rightCount[r]
		if !ok {
			rightCount[r] = 1
			continue
		}
		rightCount[r] = v + 1
	}

	sim := 0
	for _, l := range left {
		rc, ok := rightCount[l]
		if ok {
			sim = sim + (l * rc)
		}
	}

	return sim
}

func main() {
	day := aoc.Day1
	aoc.SolveExample(day, aoc.Part1, solvePart1, 11)
	aoc.SolvePuzzle(day, aoc.Part1, solvePart1, 1388114)
	aoc.SolveExample(day, aoc.Part2, solvePart2, 31)
	aoc.SolvePuzzle(day, aoc.Part2, solvePart2, 23529853)
}
