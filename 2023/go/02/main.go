package main

import (
	"advent2023/internal/aoc"
	"strconv"
	"strings"
)

type instance struct {
	blue  int
	red   int
	green int
}

type game struct {
	id        int
	instances []instance
}

const (
	blue  = "blue"
	green = "green"
	red   = "red"
)

func (i instance) valid(maxCubes instance) bool {
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

func (i instance) power() int {
	return i.green * i.red * i.blue
}

func parseInstances(input string) (instances []instance) {

	instances = []instance{}

	instanceString := strings.Split(input, ";")
	for _, iString := range instanceString {

		i := instance{}
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
		instances = append(instances, i)
	}
	return instances
}

func parseGame(input string) game {
	s := strings.TrimPrefix(input, "Game ")
	split := strings.Split(s, ":")
	id, _ := strconv.Atoi(split[0])
	instancesString := split[1]
	return game{
		id:        id,
		instances: parseInstances(instancesString),
	}
}

func solvePart1(input []string) (solution int) {
	sum := 0
	for _, line := range input {
		maxCubes := instance{
			blue:  14,
			red:   12,
			green: 13,
		}
		g := parseGame(line)
		valid := true
		for _, i := range g.instances {
			if !i.valid(maxCubes) {
				valid = false
			}
		}
		if valid {
			sum = sum + g.id
		}
	}
	return sum
}

func solvePart2(input []string) (solution int) {
	sum := 0
	for _, line := range input {
		minCubes := instance{
			blue:  0,
			red:   0,
			green: 0,
		}
		g := parseGame(line)
		for _, i := range g.instances {
			if i.blue > minCubes.blue {
				minCubes.blue = i.blue
			}
			if i.red > minCubes.red {
				minCubes.red = i.red
			}
			if i.green > minCubes.green {
				minCubes.green = i.green
			}
		}
		sum = sum + minCubes.power()
	}
	return sum
}

func main() {
	day := 2
	aoc.SolveExample(day, aoc.Part1, solvePart1, 8)
	aoc.SolvePuzzle(day, aoc.Part1, solvePart1, 2447)
	aoc.SolveExample(day, aoc.Part2, solvePart2, 2286)
	aoc.SolvePuzzle(day, aoc.Part2, solvePart2, 56322)
}
