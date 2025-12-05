package day05

import (
	"aoc2025/utils"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"
)

func Run() {
	input, err := utils.ReadInput("day05/input.txt", 5)
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
	sections := strings.Split(input, "\n\n")
	rangesRaw := strings.Split(sections[0], "\n")
	ids := strings.Split(sections[1], "\n")

	ranges := make([]utils.IntRange, 0)
	for _, r := range rangesRaw {
		parts := strings.Split(r, "-")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, utils.IntRange{First: a, Second: b})
	}

	numFresh := 0
	for _, id := range ids {
		idX, _ := strconv.Atoi(id)
		fresh := false
		for _, r := range ranges {
			if idX >= r.First && idX <= r.Second {
				fresh = true
				break
			}
		}

		if fresh {
			numFresh++
		}
	}

	return numFresh
}

func part2(input string) int {
	sections := strings.Split(input, "\n\n")
	rangesRaw := strings.Split(sections[0], "\n")
	idRanges := mergeRanges(rangesRaw)

	total := 0
	for _, r := range idRanges {
		total += r.Second - r.First + 1
	}

	return total
}

func mergeRanges(rangesRaw []string) []utils.IntRange {
	idRanges := make([]utils.IntRange, 0)

	initialIdRanges := make([]utils.IntRange, 0)
	for _, r := range rangesRaw {
		parts := strings.Split(r, "-")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		newRange := utils.NewIntRange(a, b)
		initialIdRanges = append(initialIdRanges, newRange)
	}

	slices.SortFunc(initialIdRanges, func(r1 utils.IntRange, r2 utils.IntRange) int {
		if r1.First != r2.First {
			return r1.First - r2.First
		}
		return r1.Second - r2.Second
	})

	for _, r := range initialIdRanges {
		found := false
		for i, r2 := range idRanges {
			if utils.Overlaps(r, r2) {
				found = true
				r2.First = min(r.First, r2.First)
				r2.Second = max(r.Second, r2.Second)
				idRanges[i] = r2
			}
		}

		if !found {
			idRanges = append(idRanges, r)
		}
	}

	return idRanges
}
