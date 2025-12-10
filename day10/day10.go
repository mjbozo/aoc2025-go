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

type state struct {
	joltages   []int
	presses    []int
	numPresses int
}

func (s *state) key() string {
	return fmt.Sprintf("%v", s.joltages)
}

func newState(joltagesLength int, buttonsLength int) state {
	return state{
		joltages:   make([]int, joltagesLength),
		presses:    make([]int, buttonsLength),
		numPresses: 0,
	}
}

func part2(input []string) int {
	totalPresses := 0
	totalMachines := len(input)
	fmt.Println(totalMachines)
	for m, machine := range input {
		fmt.Printf("BEGINNING MACHINE %d / %d\n", m+1, totalMachines)
		switches, joltages := parseInput(machine)

		// FIXME: NEW APPROACH
		// Use Djikstra or similar to pathfind to a solution in X-dimensional space
		// The joltages represent a point in X-dimensional space, and the buttons
		// are vector paths to reach that point

		best := make(map[string]int)
		queue := utils.BinaryHeap(func(a, b state) int {
			aScore := 0
			bScore := 0

			for i := range joltages {
				aScore += joltages[i] - a.joltages[i]
				bScore += joltages[i] - b.joltages[i]
			}

			return bScore - aScore
		})

		initialState := newState(len(joltages), len(switches))
		queue.Insert(initialState)
		best[initialState.key()] = 0
		goalState := fmt.Sprintf("%v", joltages)

		for queue.Size() > 0 {
			current, _ := queue.Pop()
			// fmt.Printf("Current node: %v, Queue size: %d\n", current, queue.Size())

			if v, ok := best[current.key()]; ok && v < current.numPresses {
				// already seen this state
				continue
			}

			if joltagesEqual(current.joltages, joltages) {
				fmt.Printf("FOUND. Remaining: %d\n", queue.Size())
				// fmt.Println(best[current.key()])
				// time.Sleep(5 * time.Second)
				break
			}

			// look for 'neighbours' in each vector direction of each button
			for i, vector := range switches {
				next := newState(len(joltages), len(switches))
				copy(next.joltages, current.joltages)
				copy(next.presses, current.presses)
				next.numPresses = current.numPresses

				next.presses[i]++
				next.numPresses++
				for _, axis := range vector {
					next.joltages[axis]++
				}

				nextValid := true
				for j := range joltages {
					if next.joltages[j] > joltages[j] {
						nextValid = false
						break
					}
				}

				if v, ok := best[next.key()]; ok {
					// already seen this state
					if v < next.numPresses {
						nextValid = false
						break
					}
				}

				if nextValid {
					best[next.key()] = next.numPresses
					queue.Insert(next)
				}
			}
		}

		totalPresses += best[goalState]
	}

	return totalPresses
}

func parseInput(machine string) ([][]int, []int) {
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
	for _, j := range joltagesStr {
		x, _ := strconv.Atoi(j)
		joltages = append(joltages, x)
	}

	return switches, joltages
}

func sum(a []int) int {
	s := 0
	for _, x := range a {
		s += x
	}
	return s
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
