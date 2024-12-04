package day1

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/underscoren/aoc2024/util"
)

func Main() {
	// read in the file as a slice of lines
	lines := util.ReadLines("day1/input")

	// transforming line text to 2D locationID int array
	locationIDs := [][]int{
		make([]int, 0, len(lines)),
		make([]int, 0, len(lines)),
	}

	for _, line := range lines {
		IDstrings := strings.Split(line, "   ")

		// convert each string to an int
		for i, s := range IDstrings {
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("Unable to convert \"%s\" to a number", s)
			}

			locationIDs[i] = append(locationIDs[i], num)
		}
	}

	slices.Sort(locationIDs[0])
	slices.Sort(locationIDs[1])

	var sum int

	for i := 0; i < len(lines); i += 1 {
		sum += util.Abs(locationIDs[0][i] - locationIDs[1][i])
	}

	fmt.Printf("Day 1 Total difference: %d \n", sum)

	var similarity int

	// sum up the number of times each number in the left list appears
	// in the right list, multiplied by itself
	for i := 0; i < len(lines); i += 1 {
		queryNumber := locationIDs[0][i]

		// find the number in the right list (binary search since it's sorted already)
		pos, found := slices.BinarySearch(locationIDs[1], queryNumber)

		if !found {
			continue
		}

		// find the total count of occurrences, by iterating and checking each element (since it's sorted)
		count := 1

		for {
			pos += 1
			if pos >= len(locationIDs[1]) {
				break
			}

			if locationIDs[1][pos] == queryNumber {
				count += 1
			} else {
				break
			}
		}

		similarity += count * queryNumber
	}

	fmt.Printf("Day 1 Similarity Score: %d\n", similarity)
}
