package day08

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
	input, err := utils.ReadInputLines("day08/input.txt", 8)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	start := time.Now()
	part1Result := part1(input, 1000)
	elapsed := time.Since(start)
	fmt.Printf("Part 1: %d (%v)\n", part1Result, elapsed)

	start = time.Now()
	part2Result := part2(input)
	elapsed = time.Since(start)
	fmt.Printf("Part 2: %d (%v)\n", part2Result, elapsed)
}

type point utils.Triple[int, int, int]

func part1(input []string, iterations int) int {
	points := make([]point, 0)

	for _, line := range input {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		points = append(points, point{x, y, z})
	}

	distances := make([]utils.Triple[int, int, int], 0)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			distances = append(distances, utils.Triple[int, int, int]{First: sqrDistance(points[i], points[j]), Second: i, Third: j})
		}
	}

	slices.SortFunc(distances, func(a, b utils.Triple[int, int, int]) int {
		return a.First - b.First
	})

	// map of point index to graph number
	graphLookup := make(map[int]int)
	graphIdx := 0

	// map of graph number to list of point indices
	revLookup := make(map[int][]int)

	for range iterations {
		shortest := distances[0]
		distances = distances[1:]

		xVal, xOk := graphLookup[shortest.Second]
		yVal, yOk := graphLookup[shortest.Third]

		// if both already exist in graphs
		if xOk && yOk {
			if xVal == yVal {
				// already part of same graph
				continue
			}

			// merge second graph into first
			secondNodes := revLookup[yVal]
			existing := revLookup[xVal]
			for _, v := range secondNodes {
				graphLookup[v] = xVal
				existing = append(existing, v)
			}
			revLookup[xVal] = existing
			delete(revLookup, yVal)
			continue
		}

		// if only one already exists in a graph, add other junction box
		if xOk {
			graphLookup[shortest.Third] = xVal
			existing := revLookup[xVal]
			existing = append(existing, shortest.Third)
			revLookup[xVal] = existing
			continue
		}

		if yOk {
			graphLookup[shortest.Second] = yVal
			existing := revLookup[yVal]
			existing = append(existing, shortest.Second)
			revLookup[yVal] = existing
			continue
		}

		// neither in existing graphs, create new one
		revLookup[graphIdx] = []int{shortest.Second, shortest.Third}
		graphLookup[shortest.Second] = graphIdx
		graphLookup[shortest.Third] = graphIdx
		graphIdx++
	}

	graphSizes := make([]int, 0)
	for _, v := range revLookup {
		graphSizes = append(graphSizes, len(v))
	}
	slices.SortFunc(graphSizes, func(a, b int) int {
		return b - a
	})

	return graphSizes[0] * graphSizes[1] * graphSizes[2]
}

func part2(input []string) int {
	numJunctions := len(input)
	points := make([]point, 0)

	for _, line := range input {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		points = append(points, point{x, y, z})
	}

	distances := make([]utils.Triple[int, int, int], 0)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			distances = append(distances, utils.Triple[int, int, int]{First: sqrDistance(points[i], points[j]), Second: i, Third: j})
		}
	}

	slices.SortFunc(distances, func(a, b utils.Triple[int, int, int]) int {
		return a.First - b.First
	})

	// map of point index to graph number
	graphLookup := make(map[int]int)
	graphIdx := 0

	// map of graph number to list of point indices
	revLookup := make(map[int][]int)

	for {
		shortest := distances[0]
		distances = distances[1:]

		xVal, xOk := graphLookup[shortest.Second]
		yVal, yOk := graphLookup[shortest.Third]

		// if both already exist in graphs
		if xOk && yOk {
			if xVal == yVal {
				// already part of same graph
				continue
			}

			// merge second graph into first
			secondNodes := revLookup[yVal]
			existing := revLookup[xVal]
			for _, v := range secondNodes {
				graphLookup[v] = xVal
				existing = append(existing, v)
			}
			revLookup[xVal] = existing
			delete(revLookup, yVal)

			if allJunctionsConnected(revLookup, numJunctions) {
				x1 := points[shortest.Second].First
				x2 := points[shortest.Third].First
				return x1 * x2
			}

			continue
		}

		// if only one already exists in a graph, add other junction box
		if xOk {
			graphLookup[shortest.Third] = xVal
			existing := revLookup[xVal]
			existing = append(existing, shortest.Third)
			revLookup[xVal] = existing

			if allJunctionsConnected(revLookup, numJunctions) {
				x1 := points[shortest.Second].First
				x2 := points[shortest.Third].First
				return x1 * x2
			}

			continue
		}

		if yOk {
			graphLookup[shortest.Second] = yVal
			existing := revLookup[yVal]
			existing = append(existing, shortest.Second)
			revLookup[yVal] = existing

			if allJunctionsConnected(revLookup, numJunctions) {
				x1 := points[shortest.Second].First
				x2 := points[shortest.Third].First
				return x1 * x2
			}

			continue
		}

		// neither in existing graphs, create new one
		revLookup[graphIdx] = []int{shortest.Second, shortest.Third}
		graphLookup[shortest.Second] = graphIdx
		graphLookup[shortest.Third] = graphIdx
		graphIdx++
	}
}

func sqrDistance(p1, p2 point) int {
	xDiff := p2.First - p1.First
	yDiff := p2.Second - p1.Second
	zDiff := p2.Third - p1.Third
	return xDiff*xDiff + yDiff*yDiff + zDiff*zDiff
}

func allJunctionsConnected(revLookup map[int][]int, numJunctions int) bool {
	if len(revLookup) == 1 {
		for _, v := range revLookup {
			if len(v) == numJunctions {
				return true
			}
		}
	}
	return false
}
