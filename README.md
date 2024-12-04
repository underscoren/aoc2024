# AoC 2024

Written in go so I can finally convince myself to learn it.

## Day 1

Simple enough. Sort, difference, sum. 

Mostly learning how to work with files and modules.

Lack of functional programming threw me for a loop, but in hindsight I should have expected that. Also, no builtin Absolute for integers? What the hell.

`slices.BinarySearch` was very nice to have though, would have been a bit annoying to implement

## Day 2

Still nothing crazy, though I went about with a lazy way to implement part 2. I'm already making my own util package for reading and parsing numbers and I've refactored Day 1 to use it.

Had to make my own "remove" function for slices, but I guess that's understandable. Still missing my functional programming paradigms.

## Day 3

This one was even easier than the above two. Match all, iterate over results. Regex to the rescue. I considered some absurdly complex lookbehind backreference pattern for Part 2 and decided I don't hate myself enough for that.

The util package grows.