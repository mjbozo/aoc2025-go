package day10

import (
	"aoc2025/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

func Run() {
	input, err := utils.ReadInputLines("day10/input.txt", 10)
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
	for _, machine := range input {
		sum += calculateButtonPresses(machine)
	}
	return sum
}

func part2(input []string) int {
	totalPresses := 0
	totalMachines := len(input)
	for q, machine := range input {
		machine = strings.ReplaceAll(machine, ") (", ")(")
		parts := strings.Split(machine, " ")
		buttons := strings.ReplaceAll(parts[1], ")(", ") (")
		buttonGroups := strings.Split(buttons, " ")
		switches := make([][]int, 0)
		for _, button := range buttonGroups {
			nums := strings.Split(strings.TrimSuffix(strings.TrimPrefix(button, "("), ")"), ",")
			group := make([]int, 0)
			for _, n := range nums {
				x, _ := strconv.Atoi(n)
				group = append(group, x)
			}
			switches = append(switches, group)
		}

		joltagesStr := strings.Split(strings.TrimSuffix(strings.TrimPrefix(parts[2], "{"), "}"), ",")

		joltages := make([]int, 0)
		upperBound := 0
		for _, j := range joltagesStr {
			x, _ := strconv.Atoi(j)
			joltages = append(joltages, x)
			upperBound += x
		}

		fmt.Printf("BEGINNING SEARCH: Machine %d / %d\n", q+1, totalMachines)
		// fmt.Printf("Joltages: %v, Switches: %v\n", joltages, switches)

		for i := 0; i <= upperBound; i++ {
			fmt.Printf("Testing with %d presses (upper bound %d)\n", i, upperBound)
			presses := make([]int, len(switches))
			if canGenerateJoltages(i, presses, switches, joltages) {
				totalPresses += i
				fmt.Printf("Machine %d requires %d presses", q+1, i)
				break
			}
		}

		// fmt.Println(switches, joltages)
		// totalPresses := 0
		// for !allJoltagesMet(joltages) {
		// 	slices.SortFunc(switches, func(a, b []int) int {
		// 		aScore := 0
		// 		bScore := 0
		//
		// 		for _, x := range a {
		// 			aScore += joltages[x]
		// 		}
		// 		for _, x := range b {
		// 			bScore += joltages[x]
		// 		}
		//
		// 		if aScore == bScore {
		// 			return len(a) - len(b)
		// 		}
		//
		// 		return bScore - aScore
		// 	})
		//
		// 	best := switches[0]
		// 	presses := math.MaxInt
		// 	for _, x := range best {
		// 		presses = min(joltages[x], presses)
		// 	}
		//
		// 	// fmt.Printf("Sorted: %v\n", switches)
		// 	for _, x := range switches[0] {
		// 		joltages[x] -= presses
		// 	}
		//
		// 	totalPresses += presses
		//
		// 	// fmt.Println(joltages, totalPresses)
		// 	// time.Sleep(2 * time.Second)
		// }
		// totalSum += totalPresses
	}

	return totalPresses
}

func canGenerateJoltages(depth int, presses []int, switches [][]int, joltages []int) bool {
	// fmt.Printf("Generated button distribution: %v\n", presses)
	if depth == 0 {
		current := make([]int, len(joltages))
		for i, press := range presses {
			for _, j := range switches[i] {
				current[j] += press
			}
		}

		return joltagesEqual(current, joltages)
	}

	for i := range switches {
		presses[i]++
		if canGenerateJoltages(depth-1, presses, switches, joltages) {
			return true
		}
		presses[i]--
	}

	return false
}

func joltagesEqual(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func allJoltagesMet(joltages []int) bool {
	for _, x := range joltages {
		if x != 0 {
			return false
		}
	}
	return true
}

func calculateButtonPresses(machine string) int {
	machine = strings.ReplaceAll(machine, ") (", ")(")
	parts := strings.Split(machine, " ")
	indicatorLights := parts[0]
	buttons := strings.ReplaceAll(parts[1], ")(", ") (")

	indicatorLights = strings.TrimSuffix(strings.TrimPrefix(indicatorLights, "["), "]")

	goal := make([]bool, 0)
	for _, c := range indicatorLights {
		if c == '#' {
			goal = append(goal, true)
		} else {
			goal = append(goal, false)
		}
	}

	buttonGroups := strings.Split(buttons, " ")
	switches := make([][]int, 0)
	for _, button := range buttonGroups {
		nums := strings.Split(strings.TrimSuffix(strings.TrimPrefix(button, "("), ")"), ",")
		group := make([]int, 0)
		for _, n := range nums {
			x, _ := strconv.Atoi(n)
			group = append(group, x)
		}
		switches = append(switches, group)
	}

	combinations := generateAllCombinations(switches)
	fewestSwitches := len(switches)
	for _, combo := range combinations {
		current := make([]bool, 0)
		for range len(goal) {
			current = append(current, false)
		}

		numSwitched := 0
		for i, s := range combo {
			if s {
				numSwitched++
				for _, buttonIdx := range switches[i] {
					current[buttonIdx] = !current[buttonIdx]
				}
			}
		}

		if stateEqual(current, goal) && numSwitched < fewestSwitches {
			fewestSwitches = numSwitched
		}
	}

	return fewestSwitches
}

func generateAllCombinations(switches [][]int) [][]bool {
	mask := int(math.Pow(2.0, float64(len(switches)))) - 1

	combinations := make([][]bool, 0)
	for mask >= 0 {
		current := mask
		combo := make([]bool, 0)

		for range len(switches) {
			combo = append(combo, false)
		}

		for i := len(combo) - 1; i >= 0; i-- {
			if current&1 == 1 {
				combo[i] = true
			}
			current = current >> 1
		}

		combinations = append(combinations, combo)
		mask--
	}

	return combinations
}

func stateEqual(a, b []bool) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
