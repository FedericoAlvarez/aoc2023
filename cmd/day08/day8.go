package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

const left = 76
const exitNode = "ZZZ"

var r = regexp.MustCompile(`(\w+)`)

type node struct {
	left, right string
}

func main() {
	part1()
	part2()

}

func part2() {
	//ring
	lines := strings.Split(input, "\n")
	directions := lines[0]
	input := make(map[string]node)

	for i := 2; i < len(lines)-1; i++ {
		allString := r.FindAllString(lines[i], -1)
		input[r.FindString(allString[0])] = node{
			left:  r.FindString(allString[1]),
			right: r.FindString(allString[2]),
		}
	}

	var nextNodes []string
	for a := range input {
		if string(a[2]) == "A" {
			nextNodes = append(nextNodes, a)
		}
	}
	var a []int
	for _, start := range nextNodes {
		a = append(a, compute(start, directions, input))
	}
	solution := LCM(a[0], a[1], a[2:]...)

	fmt.Println("Solution part 2: ", solution)
}

func compute(start string, directions string, input map[string]node) int {
	nextNode := ""
	steps := 0
	for {
		if nextNode == "" {
			nextNode = start
		}
		if directions[steps%len(directions)] == left {
			nextNode = input[nextNode].left
		} else {
			nextNode = input[nextNode].right
		}
		steps++
		if string(nextNode[2]) == "Z" && steps >= len(directions) {
			break
		}
	}
	return steps
}

func part1() {
	lines := strings.Split(input, "\n")
	directions := lines[0]
	input := make(map[string]node)

	for i := 2; i < len(lines)-1; i++ {
		allString := r.FindAllString(lines[i], -1)
		input[r.FindString(allString[0])] = node{
			left:  r.FindString(allString[1]),
			right: r.FindString(allString[2]),
		}
	}
	steps := 0
	nextNode := "AAA"
	for {
		if directions[steps%len(directions)] == left {
			nextNode = input[nextNode].left
		} else {
			nextNode = input[nextNode].right
		}
		steps++
		if nextNode == exitNode {
			break
		}
	}

	fmt.Println("Solution part 1: ", steps)
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
