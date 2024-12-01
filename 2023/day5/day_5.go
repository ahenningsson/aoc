package main

import (
	"fmt"
	"os"
	"sort"
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

	input, err := os.ReadFile("./demo_input.txt")
	check(err)
	// fmt.Printf("MyInput: %q", string(input))
	// End boiler plate
	lines := strings.Split(string(input), "\n\n")
	part1(lines)
	part2(lines)
}

type Range struct {
	destStart int64
	srcStart  int64
	rangeVal  int64
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
			if len(line) == 0 {
				continue
			}
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
			fmt.Println("seed", seed, "range slice", rangeSlice)
			seed = getMappedValue(rangeSlice, seed)
			fmt.Println("mapped val", seed)
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
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs\n")
}

// This will return the mapped value for a given source value, if it exist
func getMappedValue(ranges []Range, value int64) int64 {
	for _, rangeVal := range ranges {
		// For a given value, check the range slice if the value is in the range. If it is, return the mapped value, or else there's no reason to iterate each value in the range.
		if value >= rangeVal.srcStart && value < rangeVal.srcStart+rangeVal.rangeVal {
			return rangeVal.destStart + (value - rangeVal.srcStart)
		}
	}
	return value
}

type SeedRange struct {
	start    int64
	rangeVal int64
}

type Interval struct {
	start int64
	end   int64
}

func part2(mapBlocks []string) {
	start := time.Now()

	// Create seed ranges slices
	seedsStr := strings.TrimSpace(strings.Split(mapBlocks[0], ":")[1])
	seedsStrSlice := strings.Fields(seedsStr)

	seedRangeSlices := make([]SeedRange, len(seedsStrSlice)/2)
	index := 0
	for idx := 0; idx < len(seedsStrSlice); idx += 2 {
		startStr := seedsStrSlice[idx]
		start, err := strconv.ParseInt(startStr, 10, 64)
		check(err)

		rangeStr := seedsStrSlice[idx+1]
		rangeVal, err := strconv.ParseInt(rangeStr, 10, 64)
		check(err)

		seedRangeSlices[index] = SeedRange{start, rangeVal}
		index++
	}

	fmt.Println("seedrange slices", seedRangeSlices)

	// Create all the ranges from the input field
	almenacSlices := make([][]Range, len(mapBlocks)-1)

	for mapIdx, block := range mapBlocks {
		// This is the seeds, skip those
		if mapIdx == 0 {
			continue
		}
		// Split the block into individual lines
		lines := strings.Split(block, "\n")
		for _, line := range lines {
			if len(line) == 0 {
				continue
			}
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
			almenacSlices[mapIdx-1] = append(almenacSlices[mapIdx-1], Range{destStart, srcStart, rangeVal})
		}
	}

	// Sort accoring to the source start, so that we can iterate through the ranges in order
	for _, almenacSlice := range almenacSlices {
		sort.Slice(almenacSlice, func(i, j int) bool {
			return almenacSlice[i].srcStart < almenacSlice[j].srcStart
		})
	}

	lowestLocation := int64(1 << 60)
	for _, seedSlice := range seedRangeSlices {
		currentIntervals := []Interval{{seedSlice.start, seedSlice.start + seedSlice.rangeVal}}
		for _, rangeSlices := range almenacSlices {
			fmt.Println("seed slice", seedSlice, "range slices", rangeSlices)

			// Todo: Get the remapping. 79-93 to the range of the next destination range
		}

		// Check the intervals for the lowest value. This step happens after each mapping has been remade.
		for _, interval := range currentIntervals {
			lowestLocation = min(lowestLocation, interval.start)
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Part 2 result: ", lowestLocation)
	fmt.Println("Part 2 took: ", elapsed.Microseconds(), " μs")
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func getNewMap(seedStart int64, seedEnd int64, rangeSlices []Range) []Interval {
	intervals := []Interval{}

	// Sort the intervals according to the start value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	return intervals
}
