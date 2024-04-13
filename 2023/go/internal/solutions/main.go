package solutions

import "advent2023/internal/aoc"

var Fn map[aoc.ProblemId]aoc.Solve

func init() {
	Fn = map[aoc.ProblemId]aoc.Solve{}
	Fn[aoc.ProblemId{Day: aoc.Day1, Part: aoc.Part1}] = Day1Part1
	Fn[aoc.ProblemId{Day: aoc.Day1, Part: aoc.Part2}] = Day1Part2
}
