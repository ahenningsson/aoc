package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// check if error is nil
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Boiler plate
	err := os.Chdir("./2023/day9")
	check(err)

	input, err := os.ReadFile("./input.txt")
	check(err)
	// fmt.Printf("MyInput: %q", string(input))
	// // End boiler plate

	lines := strings.Split(string(input), "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	start := time.Now()

	// Create list of the values
	numberSlices := make([][]int64, len(lines))
	for idx, row := range lines {
		row := strings.Split(row, " ")
		// create a slice of int64 to hold the numbers
		numberList := make([]int64, len(row)+1)
		for i, val := range row {
			if len(val) == 0 {
				continue
			}
			numb, err := strconv.ParseInt(val, 10, 64)
			check(err)
			numberList[i] = numb
		}

		numberSlices[idx] = numberList
	}

	result := int64(0)
	var diffLists [][]int64
	for index, val := range numberSlices {
		diffLists = append(diffLists, numberSlices[index])
		diffList := make([]int64, len(val)-1)
		for i := 0; i < len(val)-1; i++ {
			if i == len(val)-2 {
				continue
			}
			diffList[i] = val[i+1] - val[i]
		}

		for !checkIfAllZeroes(diffList) {
			diffLists = append(diffLists, diffList)

			tempList := make([]int64, len(diffList)-1)
			for i := 0; i < len(diffList)-1; i++ {
				if i == len(diffList)-2 {
					continue
				}
				tempList[i] = diffList[i+1] - diffList[i]
			}
			diffList = tempList
		}
		diffLists = append(diffLists, diffList)

		for i := len(diffLists) - 1; i > 0; i-- {
			diffLists[i-1][len(diffLists[i-1])-1] = diffLists[i][len(diffLists[i])-1] + diffLists[i-1][len(diffLists[i-1])-2]
		}
		result += diffLists[0][len(diffLists[0])-1]

		diffLists = [][]int64{}
	}

	elapsed := time.Since(start)
	fmt.Println("Part 1 result:", result)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")
}

func part2(lines []string) {
	start := time.Now()

	// Create list of the values
	numberSlices := make([][]int64, len(lines))
	for idx, row := range lines {
		row := strings.Split(row, " ")
		// create a slice of int64 to hold the numbers
		numberList := make([]int64, len(row)+1)
		for i, val := range row {
			if len(val) == 0 {
				continue
			}
			numb, err := strconv.ParseInt(val, 10, 64)
			check(err)
			numberList[i+1] = numb
		}

		numberSlices[idx] = numberList
	}

	result := int64(0)
	var diffLists [][]int64
	for index, val := range numberSlices {
		diffLists = append(diffLists, numberSlices[index])
		diffList := make([]int64, len(val)-1)
		for i := 0; i < len(val)-1; i++ {
			if i == len(val)-2 {
				continue
			}
			diffList[i] = val[i+1] - val[i]
		}

		for !checkIfAllZeroes(diffList) {
			diffLists = append(diffLists, diffList)

			tempList := make([]int64, len(diffList)-1)
			for i := 0; i < len(diffList)-1; i++ {
				if i == len(diffList)-2 {
					continue
				}
				tempList[i] = diffList[i+1] - diffList[i]
			}
			diffList = tempList
		}
		diffLists = append(diffLists, diffList)

		for i := len(diffLists) - 1; i > 0; i-- {
			diffLists[i-1][0] = diffLists[i-1][1] - diffLists[i][0]
		}
		result += diffLists[0][0]

		diffLists = [][]int64{}
	}

	elapsed := time.Since(start)
	fmt.Println("Part 2 result:", result)
	fmt.Println("Part 2 took: ", elapsed.Microseconds(), " μs")
}

func checkIfAllZeroes(slice []int64) bool {
	for _, val := range slice {
		if val != 0 {
			return false
		}
	}
	return true
}
