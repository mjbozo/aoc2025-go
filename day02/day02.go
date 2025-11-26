package day02

	import (
		"aoc2025/utils"
		"fmt"
		"log"
		"time"
	)

	func Run() {
		input, err := utils.ReadInput("day02/input.txt", 02)
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

	func part1(input string) int {
		return 0
	}

	func part2(input string) int {
		return 0
	}