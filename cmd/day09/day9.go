package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1()
	part2()
}

func part2() {
	lines := strings.Split(input, "\n")
	solution := 0
	for _, l := range lines {
		if l == "" {
			break
		}
		values := readLineAsNumbers(l)
		end := false
		currentSequence := values
		var extrapolate []int
		extrapolate = append(extrapolate, values[0])
		for !end {
			end = true
			var nextSequence []int
			for i := 0; i < len(currentSequence)-1; i++ {
				nextSequence = append(nextSequence, diff(currentSequence, i))
			}
			extrapolate = append(extrapolate, nextSequence[0])
			currentSequence = nextSequence
			for _, aa := range nextSequence {
				if aa != 0 {
					end = false
					break
				}
			}
		}
		solution += do(extrapolate)
	}
	fmt.Println("Solution Part 2: ", solution)
}

func part1() {
	lines := strings.Split(input, "\n")
	solution := 0
	for _, l := range lines {
		if l == "" {
			break
		}
		values := readLineAsNumbers(l)
		end := false
		currentSequence := values
		extrapolate := values[len(values)-1]
		for !end {
			end = true
			var nextSequence []int
			for i := 0; i < len(currentSequence)-1; i++ {
				nextSequence = append(nextSequence, diff(currentSequence, i))
			}
			extrapolate += nextSequence[len(nextSequence)-1]
			currentSequence = nextSequence
			for _, ns := range nextSequence {
				if ns != 0 {
					end = false
					break
				}
			}
		}
		solution += extrapolate
	}
	fmt.Println("Solution Part 1: ", solution)
}

func diff(values []int, i int) int {
	x := values[i]
	y := values[i+1]
	return y - x
}

func readLineAsNumbers(l string) []int {
	a := strings.Split(l, " ")
	var r []int
	for _, aa := range a {
		r = append(r, parseString(aa))
	}
	return r
}

func parseString(s string) int {
	a, _ := strconv.Atoi(s)
	return a
}

func do(values []int) int {
	result := 0
	for i := len(values) - 1; i >= 0; i-- {
		result = values[i] - result
	}
	return result
}
