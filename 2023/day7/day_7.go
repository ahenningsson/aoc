package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// check if error is nil
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// read lines of the input file
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // append each line to the lines slice
	}
	return lines, scanner.Err()
}

func main() {

	// Boiler plate
	err := os.Chdir("./2023/day7")
	check(err)

	input, err := os.ReadFile("./demo_input.txt")
	check(err)
	// fmt.Printf("MyInput: %q", string(input))
	// // End boiler plate

	lines := strings.Split(string(input), "\n")

	part1(lines)
}

// Global constant for card values
const CARD_VAlUES = "AKQJT98765432"

func getPoint(hand string) int {
	characterOccurrences := make(map[string]int, len(CARD_VAlUES))
	for _, char := range hand {
		// fmt.Println(string(char))
		if strings.Contains(CARD_VAlUES, string(char)) {
			characterOccurrences[string(char)] += 1
		}
	}

	occurrences := make(map[int]int, 5)
	for _, value := range characterOccurrences {
		occurrences[value] += 1
	}

	// All distinct
	if occurrences[1] == 5 {
		return 1
	}
	// One pair
	if occurrences[2] == 1 {
		return 2
	}
	// Two pairs
	if occurrences[2] == 2 {
		return 3
	}
	if occurrences[3] == 1 && occurrences[1] == 2 {
		return 4
	}
	// Full house
	if occurrences[3] == 1 && occurrences[2] == 1 {
		return 5
	}
	// Four of a kind
	if occurrences[4] == 1 {
		return 6
	}
	// Straight
	if occurrences[5] == 1 {
		return 7
	}

	return 0
}

func part1(lines []string) {
	// // A map with the card values

	// // A map with the hand and its allocated points
	handPointsMap := make(map[string]int, len(lines))

	// A map with the hand and its corresponding bid value
	handBidMap := make(map[string]int, len(lines))

	// A map with the hand and its rank
	// handRankMap := make(map[string]int, len(lines))

	// Fill the bid map and the points map
	for _, row := range lines {
		rowSplit := strings.Split(row, " ")

		hand := rowSplit[0]
		bidStr := rowSplit[1]
		bid, err := strconv.ParseInt(bidStr, 10, 16)
		check(err)

		handPointsMap[hand] = getPoint(hand)
		fmt.Println("Hand: ", hand, "Point: ", getPoint(hand))

		handBidMap[hand] = int(bid)

		// TODO: send all hands with the same points to a function that will rank them
	}
}

func evaluateStrongestHand(str1 string, str2 string) {
	cardValues := map[string]int{
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

	for idx, char1 := range str1 {
		if cardValues[string(char1)] > cardValues[string(str2[idx])] {
			return
		}
	}
}
