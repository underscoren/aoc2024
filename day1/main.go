package day1

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func Part1() {
	// reading in the file
	file, err := os.Open("day1/input")
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// transforming line text to 2D locationID int array
	text := string(data)
	lines := strings.Split(strings.Trim(text, "\n"), "\n")

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
		sum += abs(locationIDs[0][i] - locationIDs[1][i])
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

		// find the total amount of occurrences
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
