package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

type point struct {
	l, c int
}

var deadEnd = point{-1, -1}

func main() {
	part1()
	part2()
}

func part2() {
	startPoint, puzzle := fillPuzzle()
	fm := firstMove(startPoint, puzzle)

	var points []point
	stop := false
	for _, actual := range fm {
		previous := startPoint
		points = append(points, previous)
		points = append(points, actual)
		for {
			if puzzle[actual.l][actual.c] == puzzle[startPoint.l][startPoint.c] {
				//"Reach starting point"
				stop = true
				break
			}
			n := next(puzzle, actual, previous)
			if n.c == -1 && n.l == -1 {
				//"Reach dead end"
				break
			}
			previous = actual
			actual = n
			points = append(points, actual)
		}
		if stop {
			break
		}
	}
	// Fisrt point is same as last point, so let's remove it
	points = points[1:]
	// shoelace formula --> https://en.wikipedia.org/wiki/Shoelace_formula
	result := 0
	for i := 0; i < len(points); i++ {
		if i == len(points)-1 {
			result += points[i].l*points[0].c - points[i].c*points[0].l
		} else {
			result += points[i].l*points[i+1].c - points[i].c*points[i+1].l
		}

	}
	area := int(math.Abs(float64(result)) / 2)
	// Pick formula --> https://en.wikipedia.org/wiki/Pick%27s_theorem
	solution := area + 1 - (len(points) / 2)
	fmt.Println("Solution part 2: ", solution)

}

func part1() {
	startPoint, puzzle := fillPuzzle()
	fm := firstMove(startPoint, puzzle)
	result := 0
	for _, actual := range fm {
		result = 0
		previous := startPoint
		for {
			result++
			if puzzle[actual.l][actual.c] == puzzle[startPoint.l][startPoint.c] {
				//"Reach starting point"
				break
			}
			n := next(puzzle, actual, previous)
			if n == deadEnd {
				// "Reach dead end"
				break
			}
			previous = actual
			actual = n
		}
	}
	fmt.Println("Solution part 1: ", result/2)
}

func fillPuzzle() (p point, puzzle [][]string) {
	lines := strings.Split(input, "\n")
	puzzle = make([][]string, len(lines))
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
			if singleItem[a] == "S" {
				p = point{i + 1, a + 1}
			}
		}
	}
	return p, addDots(puzzle, 1)
}

func firstMove(p point, puzzle [][]string) []point {
	var moves []point
	r, ok := moveRight(p, puzzle)
	if ok {
		moves = append(moves, r)
	}

	l, ok := moveLeft(p, puzzle)
	if ok {
		moves = append(moves, l)
	}

	u, ok := moveUp(p, puzzle)
	if ok {
		moves = append(moves, u)
	}

	d, ok := moveDown(p, puzzle)
	if ok {
		moves = append(moves, d)
	}
	return moves
}

func next(puzzle [][]string, actual, previous point) point {
	next := point{}
	ok := false
	switch puzzle[actual.l][actual.c] {
	case "|":
		if actual.l > previous.l {
			next, ok = moveDown(actual, puzzle)
		} else {
			next, ok = moveUp(actual, puzzle)
		}
	case "-":
		if actual.c > previous.c {
			next, ok = moveRight(actual, puzzle)
		} else {
			next, ok = moveLeft(actual, puzzle)
		}
	case "L":
		if actual.l > previous.l && actual.c == previous.c {
			next, ok = moveRight(actual, puzzle)
		} else {
			next, ok = moveUp(actual, puzzle)
		}
	case "J":
		if actual.l > previous.l {
			next, ok = moveLeft(actual, puzzle)
		} else {
			next, ok = moveUp(actual, puzzle)
		}
	case "7":
		if actual.l == previous.l {
			next, ok = moveDown(actual, puzzle)
		} else {
			next, ok = moveLeft(actual, puzzle)
		}
	case "F":
		if actual.l == previous.l {
			next, ok = moveDown(actual, puzzle)
		} else {
			next, ok = moveRight(actual, puzzle)
		}

	}
	if ok {
		return next
	}
	return deadEnd
}

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.

func moveRight(p point, puzzle [][]string) (point, bool) {
	right := point{p.l, p.c + 1}
	return right, moveRightValid(puzzle[right.l][right.c])
}
func moveLeft(p point, puzzle [][]string) (point, bool) {
	left := point{p.l, p.c - 1}
	return left, moveLeftValid(puzzle[left.l][left.c])
}
func moveUp(p point, puzzle [][]string) (point, bool) {
	up := point{p.l - 1, p.c}
	return up, moveUpValid(puzzle[up.l][up.c])
}
func moveDown(p point, puzzle [][]string) (point, bool) {
	down := point{p.l + 1, p.c}
	return down, moveDownValid(puzzle[down.l][down.c])
}

func moveDownValid(s string) bool {
	return s == "|" || s == "J" || s == "L" || s == "S"
}
func moveUpValid(s string) bool {
	return s == "|" || s == "F" || s == "7" || s == "S"
}
func moveLeftValid(s string) bool {
	return s == "-" || s == "F" || s == "L" || s == "S"
}
func moveRightValid(s string) bool {
	return s == "-" || s == "7" || s == "J" || s == "S"
}

// /Thanks ChatGPT
func addDots(matrix [][]string, extraDots int) [][]string {
	rows := len(matrix)
	cols := len(matrix[0])

	// Create a new matrix with extra dots
	newMatrix := make([][]string, rows+2)

	// Add top border
	newMatrix[0] = make([]string, cols+2)
	for i := range newMatrix[0] {
		newMatrix[0][i] = "."
	}

	// Add sides and extra dots
	for i := 0; i < rows; i++ {
		newMatrix[i+1] = append([]string{"."}, matrix[i]...)
		newMatrix[i+1] = append(newMatrix[i+1], ".")
	}

	// Add bottom border
	newMatrix[rows+1] = make([]string, cols+2)
	for i := range newMatrix[rows+1] {
		newMatrix[rows+1][i] = "."
	}

	return newMatrix
}
