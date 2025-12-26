package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func runDay(day int) error {
	dayDir := fmt.Sprintf("day%02d", day)
	
	fmt.Printf("\n%s\n", strings.Repeat("=", 60))
	fmt.Printf("Running Day %d\n", day)
	fmt.Printf("%s\n", strings.Repeat("=", 60))
	
	cmd := exec.Command("go", "run", "main.go")
	cmd.Dir = dayDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run run.go <day_number>")
		fmt.Println("Example: go run run.go 1")
		fmt.Println("         go run run.go all  (to run all days)")
		os.Exit(1)
	}
	
	arg := os.Args[1]
	
	if strings.ToLower(arg) == "all" {
		for day := 1; day <= 25; day++ {
			if err := runDay(day); err != nil {
				fmt.Printf("Error running day %d: %v\n", day, err)
			}
		}
	} else {
		day, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Error: Invalid day number")
			os.Exit(1)
		}
		if day < 1 || day > 25 {
			fmt.Println("Error: Day must be between 1 and 25")
			os.Exit(1)
		}
		if err := runDay(day); err != nil {
			fmt.Printf("Error running day %d: %v\n", day, err)
			os.Exit(1)
		}
	}
}
