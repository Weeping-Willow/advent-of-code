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

	fmt.Println("part 1:", part1(arr))
	fmt.Println("part 2:", part2(arr))
}

func part1(arr []string) int {
	total := 0
	for _, row := range arr {
		firstHalf, secondHalf := row[:len(row)/2], row[len(row)/2:]
		alreadyNoted := make([]string, 0)
		for _, c := range secondHalf {
			if !strings.Contains(firstHalf, string(c)) || strings.Contains(strings.Join(alreadyNoted, ""), string(c)) {
				continue
			}

			alreadyNoted = append(alreadyNoted, string(c))

			total += convertToPriority(c)
		}
	}

	fmt.Println(total)
	return total
}

func part2(arr []string) int {
	total := 0
	for i := 0; i < len(arr); i += 3 {
		row1, row2, row3 := arr[i], arr[i+1], arr[i+2]

		total += findCommonLetters(row1, row2, row3)
	}

	return total
}

func findCommonLetters(rows ...string) int {
	xMarksTheSpot := make(map[int32]int, 0)

	for _, row := range rows {
		alreadyNoted := make([]string, 0)
		for _, c := range row {
			if strings.Contains(strings.Join(alreadyNoted, ""), string(c)) {
				continue
			}

			alreadyNoted = append(alreadyNoted, string(c))
			val, ok := xMarksTheSpot[c]
			if !ok {
				xMarksTheSpot[c] = 0
			}

			xMarksTheSpot[c] = val + 1
		}
	}

	for key, val := range xMarksTheSpot {
		if val == len(rows) {
			return convertToPriority(key)
		}
	}

	return 0
}

func convertToPriority(val int32) int {
	if strings.ToLower(string(val)) == string(val) {
		return int(val) - 96
	}
	return int(val) - 64 + 26
}
