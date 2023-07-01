package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		splitByComma := strings.Split(row, ",")
		from1, to1, from2, to2 := getNumbers(splitByComma)

		if (from1 >= from2 && to1 <= to2) || (from1 <= from2 && to1 >= to2) {
			total++
		}
	}

	return total
}

func part2(arr []string) int {
	total := 0
	for _, row := range arr {
		splitByComma := strings.Split(row, ",")
		from1, to1, from2, to2 := getNumbers(splitByComma)

		if (from1 >= from2 && from1 <= to2) || (to1 >= from2 && to1 <= to2) || (from2 >= from1 && from2 <= to1) {
			total++
		}
	}
	return total
}

func getNumbers(splitByComma []string) (int, int, int, int) {
	column1, column2 := splitByComma[0], splitByComma[1]
	splitByDash1, splitByDash2 := strings.Split(column1, "-"), strings.Split(column2, "-")
	from1str, to1str := splitByDash1[0], splitByDash1[1]
	from2str, to2str := splitByDash2[0], splitByDash2[1]

	form1, _ := strconv.Atoi(from1str)
	to1, _ := strconv.Atoi(to1str)
	form2, _ := strconv.Atoi(from2str)
	to2, _ := strconv.Atoi(to2str)

	return form1, to1, form2, to2
}
