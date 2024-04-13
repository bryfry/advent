package main

import (
	"advent2023/internal/aoc"
	"advent2023/internal/solutions"
	"fmt"
)

func main() {
	for _, d := range aoc.ReleasedDays(2023) {

		parts := []aoc.Part{aoc.Part1, aoc.Part2}
		types := []aoc.InputType{aoc.Example, aoc.Puzzle}

		for _, p := range parts {
			for _, t := range types {
				id := aoc.ProblemId{
					Day:  d,
					Part: p,
				}
				i := aoc.Input{
					ProblemId: id,
					Type:      t,
					//Expected: expected,
				}
				var ok bool
				i.Solve, ok = solutions.Fn[id]
				if !ok {
					fmt.Printf("%s Not implemented\n", i)
					continue
				}
				fmt.Printf("%s", i)
				_ = i.Calculate()
			}
		}
	}
}
