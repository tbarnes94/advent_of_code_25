package main

import (
"fmt"
"os"
"strings"
)

// ParseInput reads and parses the input file
func ParseInput(filename string) (string, error) {
data, err := os.ReadFile(filename)
if err != nil {
return "", err
}
return strings.TrimSpace(string(data)), nil
}

// Part1 solves part 1 of the puzzle
func Part1(data string) (interface{}, error) {
// TODO: Implement solution for part 1
return nil, nil
}

// Part2 solves part 2 of the puzzle
func Part2(data string) (interface{}, error) {
// TODO: Implement solution for part 2
return nil, nil
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
