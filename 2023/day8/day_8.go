package main

import (
	"fmt"
	"os"
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
	err := os.Chdir("./2023/day8")
	check(err)

	input, err := os.ReadFile("./input.txt")
	check(err)
	// fmt.Printf("MyInput: %q", string(input))
	// // End boiler plate

	lines := strings.Split(string(input), "\n\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	start := time.Now()
	instructions := lines[0]
	sequence := strings.Split(lines[1], "\n")

	// Use the input to make something useful
	sequenceMap := make(map[string][]string, len(sequence))
	for _, val := range sequence {
		if len(val) == 0 {
			continue
		}
		split := strings.Split(val, " = ")
		mapKey := split[0]

		seqValues := strings.Split(split[1], ",")
		leftVal := strings.TrimSpace(strings.ReplaceAll(seqValues[0], "(", ""))
		rightVal := strings.TrimSpace(strings.ReplaceAll(seqValues[1], ")", ""))
		sequenceMap[mapKey] = []string{leftVal, rightVal}
	}

	const START_POINT = "AAA"
	const END_POINT = "ZZZ"

	// Find the way out
	index := 0
	nextStep := ""
	instructionLength := len(instructions)
	stepCount := 0
	for {
		currInst := instructions[stepCount%instructionLength]
		switch string(currInst) {
		case "R":
			index = 1
		case "L":
			index = 0
		}
		// handle step 1
		if stepCount == 0 {
			nextStep = sequenceMap[START_POINT][index]
		} else {
			nextStep = sequenceMap[nextStep][index]
		}

		stepCount++
		if nextStep == END_POINT {
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Part 1 result:", stepCount)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")
}

func part2(lines []string) {
	start := time.Now()
	instructions := lines[0]
	sequence := strings.Split(lines[1], "\n")

	// Use the input to make something useful
	sequenceMap := make(map[string][]string, len(sequence))
	for _, val := range sequence {
		if len(val) == 0 {
			continue
		}
		split := strings.Split(val, " = ")
		mapKey := split[0]

		seqValues := strings.Split(split[1], ",")
		leftVal := strings.TrimSpace(strings.ReplaceAll(seqValues[0], "(", ""))
		rightVal := strings.TrimSpace(strings.ReplaceAll(seqValues[1], ")", ""))
		sequenceMap[mapKey] = []string{leftVal, rightVal}
	}

	// Create a list of starting points
	startingPoints := []string{}
	for key := range sequenceMap {
		if key[2] == 'A' {
			startingPoints = append(startingPoints, key)
		}
	}

	// List of next steps
	nextStepsSlice := make([]string, len(startingPoints))
	copy(nextStepsSlice, startingPoints)

	// List of step counts/ depth count
	stepCounts := make([]int, len(startingPoints))

	// List of completed starting points
	completedStartingPoints := make([]int, len(startingPoints))
	for idx := range startingPoints {
		completedStartingPoints[idx] = 0
	}

	index := 0
	instructionLength := len(instructions)
	for {
		allReachedEndpoint := true
		for idx := range startingPoints {
			if completedStartingPoints[idx] == 1 {
				continue
			}
			currInst := instructions[stepCounts[idx]%instructionLength]
			switch string(currInst) {
			case "R":
				index = 1
			case "L":
				index = 0
			}
			nextStepsSlice[idx] = sequenceMap[nextStepsSlice[idx]][index]

			stepCounts[idx]++

			if nextStepsSlice[idx][2] == 'Z' {
				completedStartingPoints[idx] = 1
			} else {
				allReachedEndpoint = false
			}
		}

		if allReachedEndpoint {
			break
		}
	}
	// because each completion is cyclic: 11A -> 11B -> 11Z -> 11B -> 11Z (length 2, from 11A to 11Z it is length 2),
	// we just need to check the depth from each starting point, and then calculating the LCM
	lcmVal := stepCounts[0]
	for _, stepCount := range stepCounts {
		lcmVal = lcm(lcmVal, stepCount)
	}

	elapsed := time.Since(start)
	fmt.Println("Part 2 result: ", lcmVal)
	fmt.Println("Part 2 took: ", elapsed.Microseconds(), " μs")
}

// LCM algorithm
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// GCD algorithm
func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
