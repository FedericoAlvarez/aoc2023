package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

//go:embed input.txt
var input string

var numbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"eno":   "1",
	"owt":   "2",
	"eerht": "3",
	"ruof":  "4",
	"evif":  "5",
	"xis":   "6",
	"neves": "7",
	"thgie": "8",
	"enin":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

func main() {
	lines := strings.Split(input, "\n")
	total := 0
	for _, l := range lines {
		if l == "" {
			continue
		}
		r := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine|[0-9])")
		r2 := regexp.MustCompile("(eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|[0-9])")

		front := r.FindString(l)
		back := r2.FindString(Reverse(l))

		n, _ := strconv.Atoi(fmt.Sprint(numbers[front] + numbers[back]))
		total += n
	}
	fmt.Printf("Result: %v\n", total)
}

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
