package day07

import (
	"aoc2025/utils"
	"fmt"
	"log"
	"strings"
	"time"
)

func Run() {
	input, err := utils.ReadInputLines("day07/input.txt", 7)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	start := time.Now()
	part1Result := part1(input)
	elapsed := time.Since(start)
	fmt.Printf("Part 1: %d (%v)\n", part1Result, elapsed)

	start = time.Now()
	part2Result := part2(input)
	elapsed = time.Since(start)
	fmt.Printf("Part 2: %d (%v)\n", part2Result, elapsed)
}

type Pos utils.Pair[int, int]

func part1(input []string) int {
	splittersHit := utils.HashSet[Pos]{}
	startLocations := utils.HashSet[Pos]{}
	beams := make(chan Pos, len(input)*len(input[0]))

	startIndex := strings.Index(input[0], "S")
	beams <- Pos{First: startIndex, Second: 0}

	for len(beams) > 0 {
		current := <-beams

		for y := current.Second; y < len(input); y++ {
			if input[y][current.First] == '^' {
				splittersHit.Insert(Pos{First: current.First, Second: y})
				if !startLocations.Contains(Pos{First: current.First - 1, Second: y}) {
					startLocations.Insert(Pos{First: current.First - 1, Second: y})
					beams <- Pos{First: current.First - 1, Second: y}
				}
				if !startLocations.Contains(Pos{First: current.First + 1, Second: y}) {
					startLocations.Insert(Pos{First: current.First + 1, Second: y})
					beams <- Pos{First: current.First + 1, Second: y}
				}
				break
			}
		}
	}

	return len(splittersHit)
}

func part2(input []string) int {
	startPos := Pos{First: strings.Index(input[0], "S"), Second: 0}
	memo := make(map[Pos]int)
	return traverseTimeline(startPos, input, memo)
}

func traverseTimeline(startPosition Pos, input []string, memo map[Pos]int) int {
	if val, ok := memo[startPosition]; ok {
		return val
	}

	for y := startPosition.Second; y < len(input); y++ {
		if input[y][startPosition.First] == '^' {
			left := traverseTimeline(Pos{First: startPosition.First - 1, Second: y}, input, memo)
			right := traverseTimeline(Pos{First: startPosition.First + 1, Second: y}, input, memo)

			memo[Pos{First: startPosition.First - 1, Second: y}] = left
			memo[Pos{First: startPosition.First + 1, Second: y}] = right
			return left + right
		}
	}

	memo[startPosition] = 1
	return 1
}

