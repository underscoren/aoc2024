package day5

import (
	"fmt"
	"strings"

	"github.com/underscoren/aoc2024/util"
)

// check if a given list of page numbers adheres to all the rules
// returns false with the two indices where the rules were violated,
// or true and -1 for indices otherwise
func checkRules(pages []int, rules map[int][]int) (bool, int, int) {
	for i := len(pages) - 1; i >= 0; i-- {
		page := pages[i]

		// index to start checking from
		startIndex := i - 1
		if startIndex < 0 {
			continue // no point checking the first page
		}

		// list of pages which if present, must be before the current page
		pageRules := rules[page]

		// check each page before this one, if it matches any pages from the ruleset
		for j := startIndex; j >= 0; j-- {
			checkingPage := pages[j]
			for _, beforePage := range pageRules {
				if beforePage == checkingPage {
					// rule violated, this is an incorrectly ordered line
					return true, i, j
				}
			}
		}
	}

	return false, -1, -1
}

func Main() {
	// read in the input, split it into rules and pages
	text := util.ReadFile("day5/input")

	sections := strings.Split(text, "\n\n")

	ruleLines, pageLines := util.GetLines(sections[0]), util.GetLines(sections[1])

	// parse the rules into a map of page numbers to a list of pages that must come after it
	rules := make(map[int][]int)

	for _, ruleLine := range ruleLines {
		// parse strings into ints
		pages := util.AtoiFatalSlice(strings.Split(ruleLine, "|"))

		afterPage, beforePage := pages[0], pages[1]

		// ensure the map has a slice initialized
		if rules[afterPage] == nil {
			rules[afterPage] = make([]int, 0)
		}

		// add the required page to the rules for the current page
		rules[afterPage] = append(rules[afterPage], beforePage)
	}

	var middlePageSum int
	var correctOrderMiddlePageSum int // for part 2

	for _, pageLine := range pageLines {
		// parse each line of pages
		pages := util.AtoiFatalSlice(strings.Split(pageLine, ","))

		// keep track if a rule has been violated
		var ruleViolated bool

		// check each page adheres to the rules
		ruleViolated, a, b := checkRules(pages, rules)

		if !ruleViolated {
			// all pages checked, all rules satisfied
			middlePage := pages[len(pages)/2]
			middlePageSum += middlePage
		} else {
			// part 2, order them correctly
			for ruleViolated {
				// swap the positions of the offending elements
				pages[a], pages[b] = pages[b], pages[a] // i love syntactic sugar like this

				// check again that no rules have been violated
				ruleViolated, a, b = checkRules(pages, rules)
			}

			// no rules violated, now must be in order
			middlePage := pages[len(pages)/2]
			correctOrderMiddlePageSum += middlePage
		}
	}

	fmt.Printf("Day 5 Middle page sum: %d\n", middlePageSum)
	fmt.Printf("Day 5 Corrected order middle page sum: %d\n", correctOrderMiddlePageSum)
}
