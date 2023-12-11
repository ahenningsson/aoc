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
	err := os.Chdir("./2023/day5")
	check(err)

	input, err := os.ReadFile("./input.txt")
	check(err)
	// fmt.Printf("MyInput: %q", string(input))
	// // End boiler plate
	lines := strings.Split(string(input), "\n\n")
	part1(lines)
	part2(lines)
}

type Range struct {
	destStart int64
	srcStart  int64
	rangeVal  int64
}
type SeedRange struct {
	start    int64
	rangeVal int64
}

func part1(mapBlocks []string) {
	start := time.Now()

	// Create seeds slice
	seedsStr := strings.TrimSpace(strings.Split(mapBlocks[0], ":")[1])
	seeds := []int64{}
	for _, seed := range strings.Fields(seedsStr) {
		numb, err := strconv.ParseInt(seed, 10, 64)
		check(err)
		seeds = append(seeds, numb)
	}
	// Create all the ranges from the input field
	rangeSlices := make([][]Range, len(mapBlocks)-1)

	for mapIdx, block := range mapBlocks {
		// This is the seeds, skip those
		if mapIdx == 0 {
			continue
		}
		// Split the block into individual lines
		lines := strings.Split(block, "\n")
		for _, line := range lines {
			if strings.Contains(line, "map:") {
				continue
			}

			numbers := strings.Fields(line)
			// Destination start
			destStart, err := strconv.ParseInt(numbers[0], 10, 64)
			check(err)

			// Source start
			srcStart, err := strconv.ParseInt(numbers[1], 10, 64)
			check(err)

			// Range value
			rangeVal, err := strconv.ParseInt(numbers[2], 10, 64)
			check(err)

			// add the range values to their respective index.
			// Each "block" of the almanac is their respective index
			rangeSlices[mapIdx-1] = append(rangeSlices[mapIdx-1], Range{destStart, srcStart, rangeVal})
		}
	}

	lowestLocation := int64(0)
	for seedIdx, seed := range seeds {
		for idx, rangeSlice := range rangeSlices {
			seed = getMappedValue(rangeSlice, seed)
			if seedIdx == 0 && idx == len(rangeSlices)-1 {
				lowestLocation = seed
			}
			if idx == len(rangeSlices)-1 && seed < lowestLocation {
				lowestLocation = seed
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Part 1 result: ", lowestLocation)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")
}

// This will return the mapped value for a given source value, if it exist
func getMappedValue(ranges []Range, value int64) int64 {
	for _, rangeVal := range ranges {
		if value >= rangeVal.srcStart && value < rangeVal.srcStart+rangeVal.rangeVal {
			return rangeVal.destStart + (value - rangeVal.srcStart)
		}
	}
	return value
}

func part2(mapBlocks []string) {
	start := time.Now()

	// Create seed ranges slices
	seedsStr := strings.TrimSpace(strings.Split(mapBlocks[0], ":")[1])
	seedsStrSlice := strings.Fields(seedsStr)

	seedRangeSlices := make([][]SeedRange, len(seedsStrSlice)/2)
	index := 0
	for idx := 0; idx < len(seedsStrSlice); idx += 2 {
		startStr := seedsStrSlice[idx]
		start, err := strconv.ParseInt(startStr, 10, 64)
		check(err)

		rangeStr := seedsStrSlice[idx+1]
		rangeVal, err := strconv.ParseInt(rangeStr, 10, 64)
		check(err)

		seedRangeSlices[index] = append(seedRangeSlices[index], SeedRange{start, rangeVal})
		index++
	}

	// fmt.Println(seeds)
	// Create all the ranges from the input field
	rangeSlices := make([][]Range, len(mapBlocks)-1)

	for mapIdx, block := range mapBlocks {
		// This is the seeds, skip those
		if mapIdx == 0 {
			continue
		}
		// Split the block into individual lines
		lines := strings.Split(block, "\n")
		for _, line := range lines {
			if strings.Contains(line, "map:") {
				continue
			}

			numbers := strings.Fields(line)
			// Destination start
			destStart, err := strconv.ParseInt(numbers[0], 10, 64)
			check(err)

			// Source start
			srcStart, err := strconv.ParseInt(numbers[1], 10, 64)
			check(err)

			// Range value
			rangeVal, err := strconv.ParseInt(numbers[2], 10, 64)
			check(err)

			// add the range values to their respective index.
			// Each "block" of the almanac is their respective index
			rangeSlices[mapIdx-1] = append(rangeSlices[mapIdx-1], Range{destStart, srcStart, rangeVal})
		}
	}

	lowestLocation := int64(0)
	// for seedIdx, seed := range seeds {
	// 	for idx, rangeSlice := range rangeSlices {
	// 		i := sort.Search(len(rangeSlice), func(i int) bool { return rangeSlice[i].srcStart+rangeSlice[i].rangeVal > seed })

	// 		if i < len(rangeSlice) && rangeSlice[i].srcStart == seed {
	// 			seed = getMappedValue(rangeSlice, seed)
	// 		}
	// 		if seedIdx == 0 && idx == len(rangeSlices)-1 {
	// 			lowestLocation = seed
	// 		}
	// 		if seed < lowestLocation {
	// 			lowestLocation = seed
	// 		}
	// 	}
	// }

	elapsed := time.Since(start)
	fmt.Println("Part 2 result: ", lowestLocation)
	fmt.Println("Part 2 took: ", elapsed.Microseconds(), " μs")
}
