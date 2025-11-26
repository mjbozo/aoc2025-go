package daybreaker

import (
	"aoc2025/utils"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

type DaybreakError struct {
	msg string
}

func (e *DaybreakError) Error() string {
	return e.msg
}

func Create(days []string) error {
	for _, day := range days {
		dayNum := strings.TrimPrefix(day, "day")
		if len(dayNum) == 1 {
			// then i forgot to prepend day 1-9 with a zero
			day = fmt.Sprintf("day0%s", dayNum)
		}

		// validate directory does not already exist
		info, err := os.Stat(day)
		if err != nil && errors.Is(err, fs.ErrExist) {
			return err
		}

		if info != nil {
			return &DaybreakError{fmt.Sprintf("%s already exists", day)}
		}

		// create new directory with argument name from base directory
		err = os.Mkdir(day, 0750)
		if err != nil && !errors.Is(err, fs.ErrExist) {
			return err
		}

		// create `dayX.go`, `dayX_test.go`, and `dayX_example.txt` files in new directory
		err = os.WriteFile(fmt.Sprintf("%s/%s.go", day, day), fmt.Appendf(nil, `package %s

	import (
		"aoc2025/utils"
		"fmt"
		"log"
		"time"
	)

	func Run() {
		input, err := utils.ReadInput("%s/input.txt", %s)
		if err != nil {
			log.Fatalln(utils.Red(err.Error()))
		}

		start := time.Now()
		part1Result := part1(input)
		elapsed := time.Since(start)
		fmt.Printf("Part 1: %%d (%%v)\n", part1Result, elapsed)

		start = time.Now()
		part2Result := part2(input)
		elapsed = time.Since(start)
		fmt.Printf("Part 2: %%d (%%v)\n", part2Result, elapsed)
	}

	func part1(input string) int {
		return 0
	}

	func part2(input string) int {
		return 0
	}`, day, day, dayNum), 0660)

		if err != nil {
			return err
		}

		err = os.WriteFile(fmt.Sprintf("%s/%s_test.go", day, day), fmt.Appendf(nil, `package %s

	import (
		"aoc2025/utils"
		"log"
		"testing"
	)

	func TestPart1(t *testing.T) {
		input, err := utils.ReadInput("example.txt", %s)
		if err != nil {
			log.Fatalln(utils.Red(err.Error()))
		}

		expected := 0
		result := part1(input)
		if result != expected {
			t.Fatalf(utils.Red("Expected %%d, got %%d\n"), expected, result)
		}
	}

	func TestPart2(t *testing.T) {
		input, err := utils.ReadInput("example.txt", %s)
		if err != nil {
			log.Fatalln(utils.Red(err.Error()))
		}

		expected := 0
		result := part2(input)
		if result != expected {
			t.Fatalf(utils.Red("Expected %%d, got %%d\n"), expected, result)
		}
	}`, day, dayNum, dayNum), 0660)
		if err != nil {
			return err
		}

		err = os.WriteFile(fmt.Sprintf("%s/example.txt", day), []byte(""), 0660)
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", utils.Green(fmt.Sprintf("\tSuccessfully created %s files", day)))
	}

	return nil
}
