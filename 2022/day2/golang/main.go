package main

import (
	"fmt"
	"os"
	"strings"
)

type Variant struct {
	AltMe     string
	Enemy     string
	ScoreDraw int
	ScoreWin  int
	ScoreLose int
}

func main() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	cleanContent := strings.ReplaceAll(string(content), "\r", "")
	arr := strings.Split(cleanContent, "\n")
	if len(arr) == 0 {
		fmt.Println("Elf carrying that is carrying most calories is carrying: ", 0)
	}

	fmt.Println("total score at the end: ", part1(arr))
	fmt.Println("total score at the end: ", part2(arr))
}

func part1(arr []string) int {
	var allTypes = AllTypesPart1()

	totalScore := 0
	for _, c := range arr {
		both := strings.Split(c, " ")

		row1, row2 := both[0], both[1]
		item, ok := allTypes[row1]
		if !ok {
			panic("Unknown item: " + row1)
		}

		if item.AltMe == row2 { //draw
			totalScore += 3 + item.ScoreDraw
			continue
		}

		if item.Enemy == row2 { //win
			totalScore += 6 + item.ScoreLose
			continue
		}

		totalScore += item.ScoreWin
	}

	return totalScore
}

func part2(arr []string) int {
	var allTypes = AllTypesPart1()

	totalScore := 0
	for _, c := range arr {
		both := strings.Split(c, " ")

		row1, row2 := both[0], both[1]
		item, ok := allTypes[row1]
		if !ok {
			panic("Unknown item: " + row1)
		}

		if row2 == "Z" {
			totalScore += 6 + item.ScoreLose
			continue
		}

		if row2 == "Y" { //draw
			totalScore += 3 + item.ScoreDraw
			continue
		}

		totalScore += item.ScoreWin
	}

	return totalScore
}

func AllTypesPart1() map[string]Variant {
	return map[string]Variant{
		"A": {
			AltMe:     "X",
			Enemy:     "Y",
			ScoreDraw: 1,
			ScoreLose: 2,
			ScoreWin:  3,
		},
		"B": {
			AltMe:     "Y",
			Enemy:     "Z",
			ScoreWin:  1,
			ScoreDraw: 2,
			ScoreLose: 3,
		},
		"C": {
			AltMe:     "Z",
			Enemy:     "X",
			ScoreDraw: 3,
			ScoreWin:  2,
			ScoreLose: 1,
		},
	}
}
