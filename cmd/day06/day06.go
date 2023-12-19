package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// example
// input.txt
// input_part2.txt
//
//go:embed input.txt
var input string

func main() {
	//part1()
	part2()

}
func part2() {
	lines := strings.Split(input, "\n")
	r := regexp.MustCompile(`(\d+)`)

	total := 1

	time, _ := strconv.Atoi(strings.Join(r.FindAllString(lines[0], -1), ""))
	record, _ := strconv.Atoi(strings.Join(r.FindAllString(lines[1], -1), ""))

	count := 0
	for j := 0; j <= time; j++ {
		a := (time - j) * j
		if a > record {
			count++
		}
	}
	total = total * count
	//}
	fmt.Println("Result: ", total)
}

func part1() {
	lines := strings.Split(input, "\n")
	r := regexp.MustCompile(`(\d+)`)

	times := r.FindAllString(lines[0], -1)
	records := r.FindAllString(lines[1], -1)
	total := 1
	for i := 0; i < len(times); i++ {

		time, _ := strconv.Atoi(times[i])
		record, _ := strconv.Atoi(records[i])

		count := 0
		for j := 0; j <= time; j++ {
			a := (time - j) * j
			if a > record {
				//fmt.Println("Esta convinacion gana: ", j, " record: ", a)
				count++
			}
		}
		total = total * count
	}
	fmt.Println("Result: ", total)
}
