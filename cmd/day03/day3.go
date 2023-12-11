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

// var r = regexp.MustCompile("[0-9]|\\.")
var r = regexp.MustCompile("\\*")
var isNumber = regexp.MustCompile("[0-9]")

type point struct {
	i            int
	j            int
	starPosition []string
}

func main() {
	part2()
}

func part2() {
	puzzle := fillPuzzle()

	mapOfNumbersAndStars := make(map[string][]string)

	var validNumbers []string
	for i := 0; i < len(puzzle); i++ {
		numb := ""
		isValid := false
		var starPosition []string
		for j := 0; j < len(puzzle[0]); j++ {
			p := point{i: i, j: j}
			if isNumber.MatchString(puzzle[i][j]) {
				numb += puzzle[i][j]
			} else {
				if isValid {
					validNumbers = append(validNumbers, numb)
					for _, a := range removeDuplicate(starPosition) {
						mapOfNumbersAndStars[a] = append(mapOfNumbersAndStars[a], numb)
					}
				}
				numb = ""
				isValid = false
				starPosition = nil
				continue
			}
			if i == 0 {
				//first column
				if j == 0 {
					a := p.right(puzzle)
					b := p.downRight(puzzle)
					c := p.down(puzzle)
					if a || b || c {
						isValid = true
						starPosition = append(starPosition, p.starPosition...)
					}
				}
				// middle column
				if j != 0 && j != (len(puzzle[0])-1) {
					a := p.right(puzzle)
					b := p.downRight(puzzle)
					c := p.down(puzzle)
					d := p.downLeft(puzzle)
					e := p.left(puzzle)
					if a || b || c || d || e {
						isValid = true
						starPosition = append(starPosition, p.starPosition...)
					}
				}
				// last column
				if j == (len(puzzle[0]) - 1) {
					a := p.down(puzzle)
					b := p.downLeft(puzzle)
					c := p.left(puzzle)
					if a || b || c {
						isValid = true
						starPosition = append(starPosition, p.starPosition...)
					}
				}
			}
			//Middle line
			if i != 0 && i != (len(puzzle)-1) {
				if j == 0 {
					a := p.up(puzzle)
					b := p.down(puzzle)
					c := p.right(puzzle)
					d := p.upRight(puzzle)
					e := p.downRight(puzzle)
					if a || b || c || d || e {
						isValid = true
						starPosition = append(starPosition, p.starPosition...)
					}
				}
				if j != 0 && j != (len(puzzle[0])-1) {
					a := p.left(puzzle)
					b := p.right(puzzle)
					c := p.up(puzzle)
					d := p.down(puzzle)
					e := p.upLeft(puzzle)
					f := p.upRight(puzzle)
					g := p.downLeft(puzzle)
					h := p.downRight(puzzle)
					if a || b || c || d || e || f || g || h {
						isValid = true
						starPosition = append(starPosition, p.starPosition...)
					}
				}
				if j == (len(puzzle[0]) - 1) {
					a := p.up(puzzle)
					b := p.upLeft(puzzle)
					c := p.left(puzzle)
					d := p.downLeft(puzzle)
					e := p.down(puzzle)
					if a || b || c || d || e {
						isValid = true
						starPosition = append(starPosition, p.starPosition...)
					}
				}
			}
			// Bottom line
			if i == (len(puzzle) - 1) {
				if j == 0 {
					a := p.up(puzzle)
					b := p.upRight(puzzle)
					c := p.right(puzzle)
					if a || b || c {
						isValid = true
						starPosition = append(starPosition, p.starPosition...)
					}
				}
				if j != 0 && j != (len(puzzle[0])-1) {
					a := p.left(puzzle)
					b := p.upLeft(puzzle)
					c := p.up(puzzle)
					d := p.upRight(puzzle)
					e := p.right(puzzle)

					if a || b || c || d || e {
						isValid = true
						starPosition = append(starPosition, p.starPosition...)
					}
				}
				if j == (len(puzzle[0]) - 1) {
					a := p.left(puzzle)
					b := p.upLeft(puzzle)
					c := p.up(puzzle)
					if a || b || c {
						isValid = true
						starPosition = append(starPosition, p.starPosition...)
					}
				}
			}
		}
		if isValid {
			validNumbers = append(validNumbers, numb)
			for _, a := range removeDuplicate(starPosition) {
				mapOfNumbersAndStars[a] = append(mapOfNumbersAndStars[a], numb)
			}
			starPosition = nil
		}
	}
	fmt.Println(validNumbers)
	fmt.Println(mapOfNumbersAndStars)
	total := 0
	for _, value := range mapOfNumbersAndStars {
		noDuplicates := value
		if len(noDuplicates) == 2 {
			a, _ := strconv.Atoi(noDuplicates[0])
			b, _ := strconv.Atoi(noDuplicates[1])
			total += a * b
		}

	}
	fmt.Println("total: ", total)
}

//func part1() {
//	puzzle := fillPuzzle()
//
//	var validNumbers []string
//
//	for i := 0; i < len(puzzle); i++ {
//		numb := ""
//		isValid := false
//
//		for j := 0; j < len(puzzle[0]); j++ {
//			if isNumber.MatchString(puzzle[i][j]) {
//				numb += puzzle[i][j]
//			} else {
//				if isValid {
//					validNumbers = append(validNumbers, numb)
//				}
//				numb = ""
//				isValid = false
//				continue
//			}
//			//First line
//			if i == 0 {
//				//first column
//				if j == 0 {
//					if right(puzzle, i, j) || downRight(puzzle, i, j) || down(puzzle, i, j) {
//						isValid = true
//					}
//				}
//				// middle column
//				if j != 0 && j != (len(puzzle[0])-1) {
//					if right(puzzle, i, j) || downRight(puzzle, i, j) || down(puzzle, i, j) || downLeft(puzzle, i, j) || left(puzzle, i, j) {
//						isValid = true
//					}
//				}
//				// last column
//				if j == (len(puzzle[0]) - 1) {
//					if down(puzzle, i, j) || downLeft(puzzle, i, j) || left(puzzle, i, j) {
//						isValid = true
//					}
//				}
//			}
//			//Middle line
//			if i != 0 && i != (len(puzzle)-1) {
//				if j == 0 {
//					if up(puzzle, i, j) || down(puzzle, i, j) || right(puzzle, i, j) || upRight(puzzle, i, j) || downRight(puzzle, i, j) {
//						isValid = true
//					}
//				}
//				if j != 0 && j != (len(puzzle[0])-1) {
//					if left(puzzle, i, j) || right(puzzle, i, j) || up(puzzle, i, j) || down(puzzle, i, j) || upLeft(puzzle, i, j) || upRight(puzzle, i, j) || downLeft(puzzle, i, j) || downRight(puzzle, i, j) {
//						isValid = true
//					}
//				}
//				if j == (len(puzzle[0]) - 1) {
//					if up(puzzle, i, j) || upLeft(puzzle, i, j) || left(puzzle, i, j) || downLeft(puzzle, i, j) || down(puzzle, i, j) {
//						isValid = true
//					}
//				}
//			}
//			// Bottom line
//			if i == (len(puzzle) - 1) {
//				if j == 0 {
//					if up(puzzle, i, j) || upRight(puzzle, i, j) || right(puzzle, i, j) {
//						isValid = true
//					}
//				}
//				if j != 0 && j != (len(puzzle[0])-1) {
//					if left(puzzle, i, j) || upLeft(puzzle, i, j) || up(puzzle, i, j) || upRight(puzzle, i, j) || right(puzzle, i, j) {
//						isValid = true
//					}
//				}
//				if j == (len(puzzle[0]) - 1) {
//					if left(puzzle, i, j) || upLeft(puzzle, i, j) || up(puzzle, i, j) {
//						isValid = true
//					}
//				}
//			}
//		}
//		if isValid {
//			validNumbers = append(validNumbers, numb)
//		}
//	}
//
//	fmt.Println(validNumbers)
//
//	sum := 0
//	for _, n := range validNumbers {
//		atoi, _ := strconv.Atoi(n)
//		sum += atoi
//	}
//	fmt.Printf("Solution: %v\n", sum)
//}

func (p *point) right(m [][]string) bool {
	return p.checkAndSave(m, p.i, p.j+1)
}
func (p *point) left(m [][]string) bool {
	return p.checkAndSave(m, p.i, p.j-1)
}
func (p *point) up(m [][]string) bool {
	return p.checkAndSave(m, p.i-1, p.j)
}
func (p *point) down(m [][]string) bool {
	return p.checkAndSave(m, p.i+1, p.j)
}
func (p *point) upRight(m [][]string) bool {
	return p.checkAndSave(m, p.i-1, p.j+1)
}
func (p *point) upLeft(m [][]string) bool {
	return p.checkAndSave(m, p.i-1, p.j-1)
}
func (p *point) downRight(m [][]string) bool {
	return p.checkAndSave(m, p.i+1, p.j+1)
}

func (p *point) downLeft(m [][]string) bool {
	return p.checkAndSave(m, p.i+1, p.j-1)
}

func (p *point) checkAndSave(m [][]string, i, j int) bool {
	c := r.MatchString(m[i][j])
	if c {
		p.saveStar(i, j)
	}
	return c
}

func (p *point) saveStar(i, j int) {
	p.starPosition = append(p.starPosition, fmt.Sprintf("%v-%v", i, j))
}

func removeDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

//func right(p [][]string, i int, j int) bool {
//	return !r.MatchString(p[i][j+1])
//}
//func left(p [][]string, i int, j int) bool {
//	return !r.MatchString(p[i][j-1])
//}
//func up(p [][]string, i int, j int) bool {
//	return !r.MatchString(p[i-1][j])
//}
//func down(p [][]string, i int, j int) bool {
//	return !r.MatchString(p[i+1][j])
//}
//func upRight(p [][]string, i int, j int) bool {
//	return !r.MatchString(p[i-1][j+1])
//}
//func upLeft(p [][]string, i int, j int) bool {
//	return !r.MatchString(p[i-1][j-1])
//}
//func downRight(p [][]string, i int, j int) bool {
//	return !r.MatchString(p[i+1][j+1])
//}
//func downLeft(p [][]string, i int, j int) bool {
//	return !r.MatchString(p[i+1][j-1])
//}

func fillPuzzle() [][]string {
	lines := strings.Split(input, "\n")
	puzzle := make([][]string, len(lines)-1)
	for i, l := range lines {
		if l == "" {
			continue
		}
		singleItem := strings.Split(l, "")
		for a := 0; a < len(l); a++ {
			if puzzle[i] == nil {
				puzzle[i] = make([]string, len(l))
			}
			puzzle[i][a] = singleItem[a]
		}
	}
	return puzzle
}
