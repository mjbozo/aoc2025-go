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

const EPS = 1e-9
const INF = math.MaxFloat64

func part2(input []string) int {
	totalPresses := 0
	for _, machine := range input {
		switches, joltages := parseInput(machine)

		A := buildMatrix(switches, joltages)
		joltagesF := make([]float64, len(joltages))
		for i := range joltages {
			joltagesF[i] = float64(joltages[i])
		}

		presses := solve(A)
		totalPresses += presses
	}

	return totalPresses
}

type SimplexSolver struct {
	D [][]float64
	B []int
	N []int
	m int
	n int
}

func (s *SimplexSolver) pivot(r, pivotCol int) {
	k := 1.0 / s.D[r][pivotCol]

	for i := 0; i < s.m+2; i++ {
		if i == r {
			continue
		}
		for j := 0; j < s.n+2; j++ {
			if j != pivotCol {
				s.D[i][j] -= s.D[r][j] * s.D[i][pivotCol] * k
			}
		}
	}

	for i := 0; i < s.n+2; i++ {
		s.D[r][i] *= k
	}

	for i := 0; i < s.m+2; i++ {
		s.D[i][pivotCol] *= -k
	}

	s.D[r][pivotCol] = k
	s.B[r], s.N[pivotCol] = s.N[pivotCol], s.B[r]
}

func (s *SimplexSolver) find(p int) bool {
	for {
		minCol := -1
		for i := 0; i < s.n+1; i++ {
			if p == 0 && s.N[i] == -1 {
				continue
			}
			if minCol == -1 {
				minCol = i
			} else {
				if s.D[s.m+p][i] < s.D[s.m+p][minCol]-EPS ||
					(math.Abs(s.D[s.m+p][i]-s.D[s.m+p][minCol]) < EPS && s.N[i] < s.N[minCol]) {
					minCol = i
				}
			}
		}

		if s.D[s.m+p][minCol] > -EPS {
			return true
		}

		pivotCol := minCol

		minRow := -1
		for i := 0; i < s.m; i++ {
			if s.D[i][pivotCol] > EPS {
				if minRow == -1 {
					minRow = i
				} else {
					ratio1 := s.D[i][s.n+1] / s.D[i][pivotCol]
					ratio2 := s.D[minRow][s.n+1] / s.D[minRow][pivotCol]

					if ratio1 < ratio2-EPS ||
						(math.Abs(ratio1-ratio2) < EPS && s.B[i] < s.B[minRow]) {
						minRow = i
					}
				}
			}
		}

		if minRow == -1 {
			return false
		}

		s.pivot(minRow, pivotCol)
	}
}

func simplex(A [][]float64, C []float64) (float64, []float64) {
	m := len(A)
	n := len(A[0]) - 1

	N := make([]int, n+1)
	for i := range n {
		N[i] = i
	}
	N[n] = -1

	B := make([]int, m)
	for i := range m {
		B[i] = n + i
	}

	// WHY: Build tableau with proper structure
	// Columns: [original vars (n)] [artificial var (1)] [RHS (1)]
	D := make([][]float64, m+2)
	for i := range m {
		D[i] = make([]float64, n+2)
		// Copy constraint coefficients
		for j := range n {
			D[i][j] = A[i][j]
		}
		// Artificial variable coefficient = -1 for all constraints
		D[i][n] = -1
		// RHS comes from last column of A
		D[i][n+1] = A[i][n]
	}

	// Phase II objective row
	D[m] = make([]float64, n+2)
	copy(D[m], C)

	// Phase I objective row
	D[m+1] = make([]float64, n+2)
	D[m+1][n] = 1 // Minimize artificial variable

	solver := &SimplexSolver{
		D: D,
		B: B,
		N: N,
		m: m,
		n: n,
	}

	// Find most negative RHS
	r := 0
	for i := 1; i < m; i++ {
		if D[i][n+1] < D[r][n+1] {
			r = i
		}
	}

	// Phase I
	if D[r][n+1] < -EPS {
		solver.pivot(r, n)
		if !solver.find(1) {
			return -INF, nil
		}
		if D[m+1][n+1] < -EPS {
			return -INF, nil
		}
	}

	// Remove artificial variables from basis
	for i := 0; i < m; i++ {
		if solver.B[i] == -1 {
			minCol := 0
			for j := 1; j < n; j++ {
				if D[i][j] < D[i][minCol]-EPS ||
					(math.Abs(D[i][j]-D[i][minCol]) < EPS && solver.N[j] < solver.N[minCol]) {
					minCol = j
				}
			}
			solver.pivot(i, minCol)
		}
	}

	// Phase II
	if solver.find(0) {
		x := make([]float64, n)
		for i := range m {
			if solver.B[i] >= 0 && solver.B[i] < n {
				x[solver.B[i]] = D[i][n+1]
			}
		}

		objVal := 0.0
		for i := range n {
			objVal += C[i] * x[i]
		}

		return objVal, x
	}

	return -INF, nil
}

func branchAndBound(A [][]float64, n int, bval *float64) {
	C := make([]float64, n)
	for i := range n {
		C[i] = 1.0
	}

	val, x := simplex(A, C)

	if val+EPS >= *bval || val == -INF {
		return
	}

	k := -1
	v := 0
	for i := range x {
		if math.Abs(x[i]-math.Round(x[i])) > EPS {
			k = i
			v = int(x[i])
			break
		}
	}

	if k == -1 {
		if val+EPS < *bval {
			*bval = val
		}
		return
	}

	s1 := make([]float64, n+1)
	s1[k] = 1
	s1[n] = float64(v)
	newA1 := make([][]float64, len(A)+1)
	copy(newA1, A)
	newA1[len(A)] = s1
	branchAndBound(newA1, n, bval)

	s2 := make([]float64, n+1)
	s2[k] = -1
	s2[n] = float64(-(v + 1))
	newA2 := make([][]float64, len(A)+1)
	copy(newA2, A)
	newA2[len(A)] = s2
	branchAndBound(newA2, n, bval)
}

func solve(A [][]float64) int {
	n := len(A[0]) - 1
	bval := INF
	branchAndBound(A, n, &bval)
	return int(math.Round(bval))
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

func buildMatrix(switches [][]int, joltages []int) [][]float64 {
	matrix := make([][]float64, 0)
	for range len(joltages) {
		row := make([]float64, len(switches)+1)
		matrix = append(matrix, row)
	}

	for j, s := range switches {
		for _, i := range s {
			matrix[i][j] = 1
		}
	}

	for j := range joltages {
		matrix[j][len(matrix[0])-1] = float64(joltages[j])
	}

	n := len(matrix)
	for i := range n {
		row := make([]float64, 0)

		for j := range matrix[i] {
			if matrix[i][j] != 0 {
				row = append(row, -matrix[i][j])
			} else {
				row = append(row, 0)
			}
		}

		matrix = append(matrix, row)
	}

	m := len(matrix[0])
	for i := m - 2; i >= 0; i-- {
		row := make([]float64, 0)
		for j := range m {
			if i == j {
				row = append(row, -1)
			} else {
				row = append(row, 0)
			}
		}
		matrix = append(matrix, row)
	}

	return matrix
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
