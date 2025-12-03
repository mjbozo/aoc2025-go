package day03

import (
	"aoc2025/utils"
	"fmt"
	"log"
	"time"
)

func Run() {
	input, err := utils.ReadInputLines("day03/input.txt", 3)
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

func part1(input []string) int {
	sum := 0
	for _, line := range input {
		sum += maximiseBatteries(line, 2)
	}
	return sum
}

func part2(input []string) int64 {
	var sum int64
	for _, line := range input {
		sum += int64(maximiseBatteries(line, 12))
	}
	return sum
}

func maximiseBatteries(input string, height int) int {
	width := len(input)
	dp := utils.Grid[int]{}
	for range height {
		dp = append(dp, make([]int, width))
	}

	// build base cases
	dp[0][0] = int(input[0] - '0')
	for x := 1; x < width; x++ {
		dp[0][x] = max(int(input[x]-'0'), dp[0][x-1])
	}

	for y := 1; y < height; y++ {
		for x := 1; x < width; x++ {
			if y <= x {
				dp[y][x] = max(dp[y-1][x-1]*10+int(input[x]-'0'), dp[y][x-1])
			}
		}
	}

	return dp[height-1][width-1]
}
