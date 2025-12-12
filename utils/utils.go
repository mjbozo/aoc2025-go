package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strings"
)

type ResponseError struct {
	msg string
}

func (e *ResponseError) Error() string {
	return e.msg
}

// Interface for all number types, useful for generics
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Data structure for pairing two units of data
//
// First and Second fields do not need to be of the same type
type Pair[T, U any] struct {
	First  T
	Second U
}

// Default formatting for Pair datatype
//
// Returns string as "(First, Second)"
func (p Pair[T, U]) String() string {
	return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}

// Data structure for grouping three units of data
//
// First, Second, and Third fields do not need to be of the same type
type Triple[T, U, V any] struct {
	First  T
	Second U
	Third  V
}

// Default formatting for Triple datatype
//
// Returns string as "(First, Second, Third)"
func (t Triple[T, U, V]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", t.First, t.Second, t.Third)
}

// Data structure for representing a range of integers
type IntRange struct {
	First  int
	Second int
}

// IntRange constructor, ensures First is always lower bound, and Second is always upper bound
func NewIntRange(f, s int) IntRange {
	lower := f
	upper := s
	if f > s {
		lower = s
		upper = f
	}
	return IntRange{First: lower, Second: upper}
}

// Send request to AOC to retrieve input data
func RequestProblemData(day int) (string, error) {
	tokenBytes, err := os.ReadFile(".token")
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	url := fmt.Sprintf("https://adventofcode.com/2025/day/%d/input", day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Cookie", strings.TrimSpace(string(tokenBytes)))

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != 200 {
		body := make([]byte, 1024)
		res.Body.Read(body)
		return "", &ResponseError{msg: fmt.Sprintf("Request failed with status %d\nMessage: %s", res.StatusCode, string(body))}
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	input_string := strings.TrimSpace(string(bytes))
	input_string_normalised := strings.ReplaceAll(input_string, "\r\n", "\n")

	return input_string_normalised, nil
}

// Read input from specified filename, separated by newlines
func ReadInputLines(filename string, day int) ([]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			input, err := RequestProblemData(day)
			if err != nil {
				return nil, err
			}
			lines := strings.Split(input, "\n")
			return lines, nil
		} else {
			return nil, err
		}
	}

	input_string := strings.TrimSpace(string(bytes))
	input_string_normalised := strings.ReplaceAll(input_string, "\r\n", "\n")
	segments := strings.Split(input_string_normalised, "\n")

	return segments, nil
}

// Read input from specific files, separated by delim parameter
func ReadInputByDelim(filename, delim string) ([]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	input_string := strings.TrimSpace(string(bytes))
	segments := strings.Split(input_string, delim)

	return segments, nil
}

// Read input from specified filename, unseparated
func ReadInput(filename string, day int) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			lines, err := RequestProblemData(day)
			if err != nil {
				return "", err
			}
			return lines, nil
		}
		return "", err
	}

	input_string := strings.TrimSpace(string(bytes))
	return input_string, nil
}

// Transformed input into hexadecimal formatted MD5 hash
func Md5(input_string string) string {
	hash := md5.Sum([]byte(input_string))
	return hex.EncodeToString(hash[:])
}

// Calculate manhattan distance between two points
func Manhattan[T, U Number](p1, p2 Pair[T, U]) int {
	return int(math.Abs(float64(p2.Second-p1.Second)) + math.Abs(float64(p2.First-p1.First)))
}

// Calculates factorial of an integer
func Factorial(n int) int {
	factorial := 1
	for i := range n {
		factorial *= i + 1
	}
	return factorial
}

// Calculate lowest common multiple of two integers
func FindLCM(x, y int) int {
	largest := max(y, x)

	upperBound := x * y
	currentLCM := largest

	for currentLCM <= upperBound {
		if currentLCM%x == 0 && currentLCM%y == 0 {
			break
		}
		currentLCM += largest
	}

	return currentLCM
}

// Returns true if two integer ranges are overlapping
func Overlaps(r1, r2 IntRange) bool {
	return r1.First <= r2.Second && r2.First <= r1.Second
}

// Filter a slice based on a predicate function
//
// Returned slice will contain elements from original slice where predicate returned true
func Filter[S ~[]E, E any](s S, predicate func(E) bool) S {
	filtered := make([]E, 0)
	for _, v := range s {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

// Maps each element in slice according to the transformation function
func Map[S ~[]E, E any, T any](s S, transformer func(E) T) []T {
	mapped := make([]T, 0)
	for _, v := range s {
		mapped = append(mapped, transformer(v))
	}
	return mapped
}

// Abs for ints
func Abs(x int) int {
	return int(math.Abs(float64(x)))
}

// Green text colour
func Green(input string) string {
	return fmt.Sprintf("\033[32m%s\033[39m", input)
}

// Red text colour
func Red(input string) string {
	return fmt.Sprintf("\033[31m%s\033[39m", input)
}

// Blue text colour
func Blue(input string) string {
	return fmt.Sprintf("\033[34m%s\033[39m", input)
}

// Calculate determinant of matrix
func MatrixDeterminant(matrix [][]int) int {
	size := len(matrix)
	for _, row := range matrix {
		if len(row) != size {
			panic("not a square matrix")
		}
	}

	if size == 1 {
		return matrix[0][0]
	}

	if size == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	// at least 3x3
	subanswers := make([]int, 0)
	for c := range size {
		submatrix := make([][]int, size-1)
		for y := 1; y < size; y++ {
			row := make([]int, 0)
			for x := range size {
				if x != c {
					row = append(row, matrix[y][x])
				}
			}
			submatrix[y-1] = row
		}

		subanswers = append(subanswers, MatrixDeterminant(submatrix))
	}

	determinant := 0
	for i, x := range subanswers {
		if i%2 == 0 {
			determinant += matrix[0][i] * x
		} else {
			determinant -= matrix[0][i] * x
		}
	}

	return determinant
}
