package aoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	example = "e"
	puzzle  = "p"

	Part1 = Part(1)
	Part2 = Part(2)

	Day1  = Day(1)
	Day2  = Day(2)
	Day3  = Day(3)
	Day4  = Day(4)
	Day5  = Day(5)
	Day6  = Day(6)
	Day7  = Day(7)
	Day8  = Day(8)
	Day9  = Day(9)
	Day10 = Day(10)
	Day11 = Day(11)
	Day12 = Day(12)
	Day13 = Day(13)
	Day14 = Day(14)
	Day15 = Day(15)
	Day16 = Day(16)
	Day17 = Day(17)
	Day18 = Day(18)
	Day19 = Day(19)
	Day20 = Day(20)
	Day21 = Day(21)
	Day22 = Day(22)
	Day23 = Day(23)
	Day24 = Day(24)
	Day25 = Day(25)

	UnknownExpected = -1
)

type InputType string
type Day int
type Part int

type solve func([]string) int

type Input struct {
	Day      Day
	Part     Part
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

func SolveExample(day Day, part Part, fn solve, expected int) {
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

func SolvePuzzle(day Day, part Part, fn solve, expected int) {
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
