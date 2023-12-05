package main

import (
	_ "embed"
	"fmt"
	"regexp"
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
	r2 := regexp.MustCompile(`\d+`)

	r3 := regexp.MustCompile(`(\d+) red`)
	r4 := regexp.MustCompile(`(\d+) green`)
	r5 := regexp.MustCompile(`(\d+) blue`)

	total := 0
	maxColor := func(color []string) int {
		m := 0
		for _, c := range color {
			value, _ := strconv.Atoi(r2.FindString(c))
			if m == 0 || m < value {
				m = value
			}
		}
		return m
	}

	for _, l := range lines {
		if l == "" {
			continue
		}
		red := r3.FindAllString(l, -1)
		green := r4.FindAllString(l, -1)
		blue := r5.FindAllString(l, -1)

		total += maxColor(red) * maxColor(green) * maxColor(blue)

	}
	fmt.Printf("Solution: %v \n", total)
}

func part1() {
	r1 := regexp.MustCompile(`Game (\d+):`)
	r2 := regexp.MustCompile(`\d+`)

	r3 := regexp.MustCompile(`(\d+) red`)
	r4 := regexp.MustCompile(`(\d+) green`)
	r5 := regexp.MustCompile(`(\d+) blue`)

	total := 0
	validColor := func(color []string, max int) bool {
		for _, r := range color {
			value, _ := strconv.Atoi(r2.FindString(r))
			if value > max {
				return false
			}
		}
		return true
	}
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		if l == "" {
			continue
		}
		matches := r1.FindString(l)
		gameNumber := r2.FindString(matches)

		red := r3.FindAllString(l, -1)
		green := r4.FindAllString(l, -1)
		blue := r5.FindAllString(l, -1)
		validGame := true

		validGame = validColor(red, 12) && validColor(green, 13) && validColor(blue, 14)

		if validGame {
			n, _ := strconv.Atoi(gameNumber)
			total += n
		}

	}
	fmt.Printf("Solution: %v \n", total)
}
