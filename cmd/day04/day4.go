package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	//part1()
	part2()
}

func part1() {
	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`:`)
	re2 := regexp.MustCompile(`\|`)
	re3 := regexp.MustCompile(`\d+`)
	var result float64

	for _, l := range lines {
		if l == "" {
			continue
		}
		matches := re.Split(l, -1)
		matches2 := re2.Split(matches[1], -1)
		set1 := strings.TrimSpace(matches2[0])
		set2 := strings.TrimSpace(matches2[1])

		firstGroup := re3.FindAllString(set1, -1)
		secondGroup := re3.FindAllString(set2, -1)

		m := make(map[string]int)
		for _, i := range firstGroup {
			m[i] = 1
		}
		for _, i := range secondGroup {
			m[i] = m[i] + 1
		}

		total := 0
		for _, value := range m {
			if value > 1 {
				total++
			}
		}

		if total > 0 {
			result += math.Pow(2, float64(total-1))
		}
	}
	fmt.Println(result)
}

func part2() {
	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`:`)
	re2 := regexp.MustCompile(`\|`)
	re3 := regexp.MustCompile(`\d+`)

	var result []int

	for _, l := range lines {
		if l == "" {
			continue
		}
		matches := re.Split(l, -1)
		matches2 := re2.Split(matches[1], -1)
		set1 := strings.TrimSpace(matches2[0])
		set2 := strings.TrimSpace(matches2[1])

		firstGroup := re3.FindAllString(set1, -1)
		secondGroup := re3.FindAllString(set2, -1)

		m := make(map[string]int)
		for _, i := range firstGroup {
			m[i] = 1
		}
		for _, i := range secondGroup {
			m[i] = m[i] + 1
		}

		total := 0
		for _, value := range m {
			if value > 1 {
				total++
			}
		}
		result = append(result, total)
	}
	total := 0
	for index, value := range result {
		if value == 0 {
			total++
			continue
		}
		total += score(result[index:], 1)

	}
	fmt.Println("Total: ", total)
}

func score(input []int, count int) int {
	a := input[0]
	if a == 0 {
		return count
	}
	for i := 0; i < a; i++ {
		count++
		count = score(input[i+1:], count)
	}
	return count
}
