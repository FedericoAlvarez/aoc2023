package main

import (
	_ "embed"
	"fmt"
	"math/big"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// input.txt
// example.txt
// test.txt

//go:embed input.txt
var input string

var cards = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}
var cardsPart2 = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

var handsScore = map[string]int{
	"5":         7,
	"4-1":       6,
	"3-2":       5,
	"3-1-1":     4,
	"2-2-1":     3,
	"2-1-1-1":   2,
	"1-1-1-1-1": 1,
}

type hand struct {
	cards string
	bet   string
}

func (h hand) betToInt() int {
	a, _ := strconv.Atoi(h.bet)
	return a
}

func main() {
	part1()
	part2()
}

func part2() {
	lines := strings.Split(input, "\n")
	partialScore := make(map[int][]hand)

	for _, l := range lines {
		if l == "" {
			continue
		}
		split := strings.Split(l, " ")
		currentHand := hand{cards: split[0], bet: split[1]}
		var handScore []int
		differentCards := 0
		for k := range cardsPart2 {
			count := strings.Count(currentHand.cards, k)
			differentCards += count
			if count > 0 {
				handScore = append(handScore, count)
			}
		}
		sort.Sort(sort.Reverse(sort.IntSlice(handScore)))
		if differentCards != 5 {
			if len(handScore) == 0 {
				handScore = append(handScore, 5)
			} else {
				handScore[0] += 5 - differentCards
			}
		}
		formattedHand := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(handScore)), "-"), "[]")
		score := handsScore[formattedHand]
		partialScore[score] = append(partialScore[score], currentHand)
	}

	score := big.NewInt(0)
	rank := len(lines) - 1
	for i := 7; i >= 0; i-- {
		if partialScore[i] != nil {
			sort.Sort(handSorterPart2(partialScore[i]))
			slices.Reverse(partialScore[i])
			for _, s := range partialScore[i] {
				score.Add(score, big.NewInt(int64(rank*s.betToInt())))
				rank--
			}
		}

	}

	fmt.Println("Solution part 2 :", score)
}

func part1() {
	lines := strings.Split(input, "\n")

	partialScore := make(map[int][]hand)

	for _, l := range lines {
		if l == "" {
			continue
		}
		split := strings.Split(l, " ")
		currentHand := hand{cards: split[0], bet: split[1]}
		var handScore []int
		for k := range cards {
			count := strings.Count(currentHand.cards, k)
			if count > 0 {
				handScore = append(handScore, count)
			}
		}
		sort.Sort(sort.Reverse(sort.IntSlice(handScore)))
		formattedHand := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(handScore)), "-"), "[]")
		score := handsScore[formattedHand]
		partialScore[score] = append(partialScore[score], currentHand)
	}

	score := big.NewInt(0)
	rank := len(lines) - 1
	for i := 7; i >= 0; i-- {
		if partialScore[i] != nil {
			sort.Sort(handSorter(partialScore[i]))
			slices.Reverse(partialScore[i])
			for _, s := range partialScore[i] {
				score.Add(score, big.NewInt(int64(rank*s.betToInt())))
				rank--
			}
		}

	}
	fmt.Println("Solution part 1:", score)
}

type handSorter []hand

func (h handSorter) Len() int {
	return len(h)
}
func (h handSorter) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h handSorter) Less(i, j int) bool {
	for a := 0; a < len(h[i].cards); a++ {
		if cards[string(h[i].cards[a])] == cards[string(h[j].cards[a])] {
			continue
		}
		if cards[string(h[i].cards[a])] < cards[string(h[j].cards[a])] {
			return true
		} else {
			return false
		}
	}
	return false
}

type handSorterPart2 []hand

func (h handSorterPart2) Len() int {
	return len(h)
}
func (h handSorterPart2) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h handSorterPart2) Less(i, j int) bool {
	for a := 0; a < len(h[i].cards); a++ {
		if cardsPart2[string(h[i].cards[a])] == cardsPart2[string(h[j].cards[a])] {
			continue
		}
		if cardsPart2[string(h[i].cards[a])] < cardsPart2[string(h[j].cards[a])] {
			return true
		} else {
			return false
		}
	}
	return false
}
