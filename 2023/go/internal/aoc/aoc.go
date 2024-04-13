package aoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	Example = InputType("e")
	Puzzle  = InputType("p")

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

type Solve func([]string) int

type ProblemId struct {
	Day  Day
	Part Part
}

type Input struct {
	ProblemId
	Solve
	Type     InputType
	Expected int
}

func (i Input) String() string {
	var t string = "Example"
	if i.Type == Puzzle {
		t = "Puzzle"
	}
	return fmt.Sprintf("Day %02d, Part %d %7s: ", i.Day, i.Part, t)

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
		fmt.Printf("%8d\n", solution)
		return nil
	}
	if solution != i.Expected {
		fmt.Printf("%8d ❌\n", solution)
		return nil
	}
	fmt.Printf("%8d ✅\n", solution)
	return nil
}

func EventDay(year int) (day int) {

	start := time.Date(year, time.December, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	end := time.Date(year, time.December, 25, 0, 0, 0, 0, time.UTC)

	if now.After(end) {
		return 25
	}
	if now.Before(start) {
		return 0
	}
	return now.Day()
}

func ReleasedDays(year int) (days []Day) {
	today := EventDay(year)
	days = []Day{}
	for d := 1; d <= today; d++ {
		days = append(days, Day(d))
	}
	return days
}

func SolveExample(day Day, part Part, fn Solve, expected int) {
	i := Input{
		ProblemId: ProblemId{
			Day:  day,
			Part: part,
		},
		Type:     Example,
		Solve:    fn,
		Expected: expected,
	}
	err := i.Calculate()
	if err != nil {
		log.Fatal(err)
	}
}

func SolvePuzzle(day Day, part Part, fn Solve, expected int) {
	i := Input{
		ProblemId: ProblemId{
			Day:  day,
			Part: part,
		},
		Type:     Puzzle,
		Solve:    fn,
		Expected: expected,
	}
	err := i.Calculate()
	if err != nil {
		log.Fatal(err)
	}
}
