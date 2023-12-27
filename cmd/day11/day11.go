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
	col, row int
}

func main() {
	part1()
	part2()
}

func part2() {
	puzzle := fillPuzzle()
	puzzle = expandRowsPart2(puzzle)
	puzzle = transpose(puzzle)
	puzzle = expandRowsPart2(puzzle)
	puzzle = transpose(puzzle)
	// printPuzzle(puzzle)

	var galaxies []point
	rows := len(puzzle)
	cols := len(puzzle[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if puzzle[i][j] == "#" {
				galaxies = append(galaxies, point{i, j})
			}
		}
	}

	solution := 0
	for i := 0; i < len(galaxies); i++ {
		nextGalaxies := galaxies[i+1:]
		for j := 0; j < len(nextGalaxies); j++ {
			galaxy := galaxies[i]
			galaxyToVisit := nextGalaxies[j]
			expansion := 0
			expansionMultiplier := 1000000 - 1
			if galaxy.row < galaxyToVisit.row {
				for x := galaxy.row; x < galaxyToVisit.row; x++ {
					if puzzle[0][x] == "X" {
						expansion += expansionMultiplier
					}
				}
			} else {
				for x := galaxyToVisit.row; x < galaxy.row; x++ {
					if puzzle[0][x] == "X" {
						expansion += expansionMultiplier
					}
				}
			}
			if galaxy.col < galaxyToVisit.col {
				for x := galaxy.col; x < galaxyToVisit.col; x++ {
					if puzzle[x][0] == "X" {
						expansion += expansionMultiplier
					}
				}
			} else {
				for x := galaxyToVisit.col; x < galaxy.col; x++ {
					if puzzle[x][0] == "X" {
						expansion += expansionMultiplier
					}
				}
			}
			solution += int(math.Abs(float64(galaxy.col-galaxyToVisit.col))) + int(math.Abs(float64(galaxy.row-galaxyToVisit.row))) + expansion
		}
	}
	fmt.Println("Solution part 2: ", solution)
}
func part1() {

	puzzle := fillPuzzle()
	puzzle = expandRows(puzzle)
	puzzle = transpose(puzzle)
	puzzle = expandRows(puzzle)
	puzzle = transpose(puzzle)
	// printPuzzle(puzzle)

	var galaxies []point
	rows := len(puzzle)
	cols := len(puzzle[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if puzzle[i][j] == "#" {
				galaxies = append(galaxies, point{i, j})
			}
		}
	}

	solution := 0
	for i := 0; i < len(galaxies); i++ {
		nextGalaxies := galaxies[i+1:]
		for j := 0; j < len(nextGalaxies); j++ {
			a := galaxies[i]
			b := nextGalaxies[j]
			solution += int(math.Abs(float64(a.col-b.col))) + int(math.Abs(float64(a.row-b.row)))
		}
	}
	fmt.Println("Solution part 1: ", solution)
}

func expandRows(puzzle [][]string) [][]string {
	rows := len(puzzle)
	cols := len(puzzle[0])
	var expandedPuzzle [][]string
	for i := 0; i < rows; i++ {
		expand := true
		for j := 0; j < cols; j++ {
			if puzzle[i][j] != "." {
				expand = false
			}
		}
		expandedPuzzle = append(expandedPuzzle, puzzle[i])
		if expand {
			expandedPuzzle = append(expandedPuzzle, createDotLine(cols))
		}
	}
	return expandedPuzzle
}
func expandRowsPart2(puzzle [][]string) [][]string {
	rows := len(puzzle)
	cols := len(puzzle[0])
	var expandedPuzzle [][]string
	for i := 0; i < rows; i++ {
		expand := true
		for j := 0; j < cols; j++ {
			if puzzle[i][j] == "#" {
				expand = false
				break
			}
		}

		if expand {
			expandedPuzzle = append(expandedPuzzle, createCrossLine(cols))
		} else {
			expandedPuzzle = append(expandedPuzzle, puzzle[i])
		}
	}
	return expandedPuzzle
}

func createDotLine(l int) []string {
	newLine := make([]string, l)
	for i := 0; i < l; i++ {
		newLine[i] = "."
	}
	return newLine
}
func createCrossLine(l int) []string {
	newLine := make([]string, l)
	for i := 0; i < l; i++ {
		newLine[i] = "X"
	}
	return newLine
}

func fillPuzzle() [][]string {
	lines := strings.Split(input, "\n")
	puzzle := make([][]string, len(lines))
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

func transpose(matrix [][]string) [][]string {
	rows := len(matrix)
	cols := len(matrix[0])

	// Create a new matrix with swapped rows and columns
	transposed := make([][]string, cols)
	for i := range transposed {
		transposed[i] = make([]string, rows)
		for j := range transposed[i] {
			transposed[i][j] = matrix[j][i]
		}
	}

	return transposed
}

func printPuzzle(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(i, " ", s[i])
	}
}
