package day11

import (
	"aoc2025/utils"
	"fmt"
	"log"
	"strings"
	"time"
)

func Run() {
	input, err := utils.ReadInputLines("day11/input.txt", 11)
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
	paths := 0
	connections := make(map[string][]string)

	for _, line := range input {
		parts := strings.Split(line, ": ")
		source := parts[0]
		destinations := strings.Split(parts[1], " ")
		connections[source] = destinations
	}

	traverse("you", connections, &paths)
	return paths
}

func traverse(current string, connections map[string][]string, paths *int) {
	if current == "out" {
		*paths++
		return
	}

	for _, connection := range connections[current] {
		traverse(connection, connections, paths)
	}
}

func part2(input []string) int {
	connections := make(map[string][]string)

	for _, line := range input {
		parts := strings.Split(line, ": ")
		source := parts[0]
		destinations := strings.Split(parts[1], " ")
		connections[source] = destinations
	}

	svrToDAC := 0
	svrToFFT := 0
	DACToFFT := 0
	FFTToDAC := 0
	DACToOut := 0
	FFTToOut := 0

	memo := make(map[string]int)
	traverseTo("svr", "dac", connections, &svrToDAC, &memo)

	memo = make(map[string]int)
	traverseTo("svr", "fft", connections, &svrToFFT, &memo)

	memo = make(map[string]int)
	traverseTo("dac", "fft", connections, &DACToFFT, &memo)

	memo = make(map[string]int)
	traverseTo("fft", "dac", connections, &FFTToDAC, &memo)

	memo = make(map[string]int)
	traverseTo("dac", "out", connections, &DACToOut, &memo)

	memo = make(map[string]int)
	traverseTo("fft", "out", connections, &FFTToOut, &memo)

	return (svrToDAC * DACToFFT * FFTToOut) + (svrToFFT * FFTToDAC * DACToOut)
}

func traverseTo(current string, target string, connections map[string][]string, paths *int, memo *map[string]int) int {
	if current == target {
		*paths++
		return 1
	}

	if val, ok := (*memo)[current]; ok {
		*paths += val
		return val
	}

	total := 0
	for _, connection := range connections[current] {
		total += traverseTo(connection, target, connections, paths, memo)
	}
	(*memo)[current] = total
	return total
}
