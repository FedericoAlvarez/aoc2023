package main

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

func TestParsing(t *testing.T) {
	s := "two1nine2oneight"

	total := 0
	r := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine|[0-9])")
	r2 := regexp.MustCompile("(eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|[0-9])")

	front := r.FindString(s)
	back := r2.FindString(Reverse(s))

	fmt.Println("front: " + front)
	fmt.Println("back: " + back)

	n, err := strconv.Atoi(fmt.Sprint(numbers[front] + numbers[back]))
	if err != nil {
		panic(err)
	}
	total += n
	fmt.Println(total)
}
