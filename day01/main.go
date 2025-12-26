package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ParseInput reads and parses the input file
func ParseInput(filename string) ([]int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	numbers := make([]int, 0, len(lines))
	
	for _, line := range lines {
		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	
	return numbers, nil
}

// Part1 solves part 1 of the puzzle - sum of all numbers
func Part1(numbers []int) (interface{}, error) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum, nil
}

// Part2 solves part 2 of the puzzle - product of all numbers
func Part2(numbers []int) (interface{}, error) {
	product := 1
	for _, num := range numbers {
		product *= num
	}
	return product, nil
}

func main() {
	// Test with example input
	exampleData, err := ParseInput("example.txt")
	if err != nil {
		fmt.Printf("Error reading example.txt: %v\n", err)
	} else {
		result1, _ := Part1(exampleData)
		fmt.Printf("Example Part 1: %v\n", result1)
		result2, _ := Part2(exampleData)
		fmt.Printf("Example Part 2: %v\n", result2)
	}

	fmt.Println()

	// Solve with real input
	data, err := ParseInput("input.txt")
	if err != nil {
		fmt.Printf("Error reading input.txt: %v\n", err)
		os.Exit(1)
	}

	result1, err := Part1(data)
	if err != nil {
		fmt.Printf("Error in Part 1: %v\n", err)
	} else {
		fmt.Printf("Part 1: %v\n", result1)
	}

	result2, err := Part2(data)
	if err != nil {
		fmt.Printf("Error in Part 2: %v\n", err)
	} else {
		fmt.Printf("Part 2: %v\n", result2)
	}
}
