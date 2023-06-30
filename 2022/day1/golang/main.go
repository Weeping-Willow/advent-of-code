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

	currentElfCalories := 0
	highestElfCalories := make([]int, 3, 3)
	for _, c := range arr {
		if c == "" {
			highestElfCalories = findTopEaters(currentElfCalories, highestElfCalories)
			currentElfCalories = 0
			continue
		}

		currentElfCalories += convertElfCalories(c)
	}

	fmt.Println("Elf carrying that is carrying most calories is carrying: ", highestElfCalories[0])
	fmt.Println("Top 3 elfs are carrying in total: ", total(highestElfCalories))
}

func findTopEaters(current int, highest []int) []int {
	for i := 0; i < len(highest); i++ {
		if current > highest[i] {
			temp := highest[i]
			highest[i] = current

			return findTopEaters(temp, highest)
		}
	}
	return highest
}

func convertElfCalories(calories string) int {
	val, err := strconv.Atoi(calories)
	if err != nil {
		fmt.Println(calories, err)
		return 0
	}

	return val
}

func total(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}

	return total
}
