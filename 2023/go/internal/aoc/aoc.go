package aoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	example         = "e"
	puzzle          = "p"
	Part1           = 1
	Part2           = 2
	UnknownExpected = -1
)

type InputType string

type solve func([]string) int

type Input struct {
	Day      int
	Part     int
	Type     InputType
	Solve    solve
	Expected int
}

func (i Input) Calculate() (err error) {

	inputFilename := fmt.Sprintf("../input/%02d-%s%d.txt", i.Day, i.Type, i.Part)
	input := []string{}
	file, err := os.Open(inputFilename)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", inputFilename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan by lines: %w", err)
	}

	solution := i.Solve(input)
	if i.Expected == UnknownExpected {
		fmt.Printf("%-6d\n", solution)
		return nil
	}
	if solution != i.Expected {
		fmt.Printf("%-6d ❌\n", solution)
		return nil
	}
	fmt.Printf("%-6d ✅\n", solution)
	return nil
}

func SolveExample(day int, part int, fn solve, expected int) {
	fmt.Printf("Day %02d, Part %d Example: ", day, part)
	i := Input{
		Day:      day,
		Part:     part,
		Type:     example,
		Solve:    fn,
		Expected: expected,
	}
	err := i.Calculate()
	if err != nil {
		log.Fatal(err)
	}
}

func SolvePuzzle(day int, part int, fn solve, expected int) {
	fmt.Printf("Day %02d, Part %d Puzzle:  ", day, part)
	i := Input{
		Day:      day,
		Part:     part,
		Type:     puzzle,
		Solve:    fn,
		Expected: expected,
	}
	err := i.Calculate()
	if err != nil {
		log.Fatal(err)
	}
}
