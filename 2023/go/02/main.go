package main

import (
	"advent2023/internal/aoc"
	"strconv"
	"strings"
)

// A collection of cubes
// Can be used to represent:
// - A constraint (e.g. maxCubes)
// - A collected state (e.g. minCubes)
// - An instance of a revield set of cubes
type bag struct {
	blue  int
	red   int
	green int
}

// Game stuct containing mutliple instances of 
// subsets of cubes that were revealed from the bag
type game struct {
	id   int
	bags []bag
}

const (
	blue  = "blue"
	green = "green"
	red   = "red"
)

func (i bag) valid(maxCubes bag) bool {
	if i.blue > maxCubes.blue {
		return false
	}
	if i.green > maxCubes.green {
		return false
	}
	if i.red > maxCubes.red {
		return false
	}
	return true
}

func (i bag) power() int {
	return i.green * i.red * i.blue
}

// Example input:
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
func parseGame(input string) game {
	input = strings.TrimPrefix(input, "Game ")
	idSplit := strings.Split(input, ":")
	id, _ := strconv.Atoi(idSplit[0])

	bags := []bag{}
	bagString := strings.Split(idSplit[1], ";")
	for _, iString := range bagString {

		i := bag{}
		cubes := strings.Split(iString, ",")
		for _, c := range cubes {

			c = strings.Trim(c, " ")
			countSplit := strings.Split(c, " ")
			if countSplit[1] == blue {
				i.blue, _ = strconv.Atoi(countSplit[0])
			}
			if countSplit[1] == red {
				i.red, _ = strconv.Atoi(countSplit[0])
			}
			if countSplit[1] == green {
				i.green, _ = strconv.Atoi(countSplit[0])
			}
		}
		bags = append(bags, i)
	}

	return game{
		id:   id,
		bags: bags,
	}
}

func solvePart1(input []string) (solution int) {
	solution = 0
	maxCubes := bag{
		blue:  14,
		red:   12,
		green: 13,
	}
	for _, line := range input {
		game := parseGame(line)
		valid := true
		for _, bag := range game.bags {
			if !bag.valid(maxCubes) {
				valid = false
				break
			}
		}
		if valid {
			solution += game.id
		}
	}
	return solution
}

// Show that you can maintain state across bags in a game
// Find the minimum number of cubes of each color across bags
func solvePart2(input []string) (solution int) {
	solution = 0
	for _, line := range input {
		minCubes := bag{
			blue:  0,
			red:   0,
			green: 0,
		}
		g := parseGame(line)
		for _, bag := range g.bags {
			if bag.blue > minCubes.blue {
				minCubes.blue = bag.blue
			}
			if bag.red > minCubes.red {
				minCubes.red = bag.red
			}
			if bag.green > minCubes.green {
				minCubes.green = bag.green
			}
		}
		solution += minCubes.power()
	}
	return solution
}

func main() {
	day := aoc.Day2
	aoc.SolveExample(day, aoc.Part1, solvePart1, 8)
	aoc.SolvePuzzle(day, aoc.Part1, solvePart1, 2447)
	aoc.SolveExample(day, aoc.Part2, solvePart2, 2286)
	aoc.SolvePuzzle(day, aoc.Part2, solvePart2, 56322)
}
