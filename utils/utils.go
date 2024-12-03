package utils

import (
	"bufio"
	"os"
)

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

// boiler plate
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	CheckErr(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // append each line to the lines slice
	}
	return lines, scanner.Err()
}
