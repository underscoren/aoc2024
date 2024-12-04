package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/underscoren/aoc2024/util"
)

func stringToInts(line string) []int {
	numbers := make([]int, 0, 4)

	parts := strings.Split(line, " ")
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, num)
	}

	return numbers
}

// returns true when the numbers in the slice are all increasing
// or all decreasing between 1 and 3
func isSafe(numbers []int) bool {
	// create slice of ints with difference between the numbers
	deltas := make([]int, len(numbers)-1)

	for i := 0; i < (len(numbers) - 1); i++ {
		deltas[i] = numbers[i] - numbers[i+1]
	}

	// the first delta should match all the others
	target := deltas[0] < 0

	for _, delta := range deltas {
		// check that all are increasing or all decreasing
		if (delta < 0) != target {
			return false
		}

		// check that all are within the correct range
		absDelta := util.Abs(delta)
		if absDelta < 1 || absDelta > 3 {
			return false
		}
	}

	return true
}

func Main() {
	lines := util.ReadLines("day2/input")

	readings := make([][]int, len(lines))

	// count how many lines are "safe" ->
	// all numbers in line decreasing or all increasing by 1-3
	var safe int

	for i, line := range lines {
		// convert line to slice of ints
		numbers := stringToInts(line)
		readings[i] = numbers

		lineSafe := isSafe(numbers)

		if lineSafe {
			safe += 1
		}
	}

	fmt.Printf("Day 2 safe count: %d\n", safe)

	var safeDampened int

	// account for 1 bad reading
	for _, numbers := range readings {
		lineSafe := isSafe(numbers)

		if !lineSafe {
			var dampenedSafe bool

			// try again with each number removed
			for pos := range numbers {
				if isSafe(util.Remove(numbers, pos)) {
					dampenedSafe = true
					break
				}
			}

			// all attempts failed, don't count this line
			if !dampenedSafe {
				continue
			}
		}

		safeDampened++
	}

	fmt.Printf("Day 2 safe dampened: %d\n", safeDampened)
}
