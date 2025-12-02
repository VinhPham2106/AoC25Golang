package main

import (
	"fmt"
	"strings"

	"github.com/vinhpham2106/aoc25golang/utils"
)

func main() {
	day := 1 // Example day
	input, err := utils.GetInput(day)
	if err != nil {
		fmt.Printf("Error fetching input: %v\n", err)
		return
	}
	// DEBUG
	// fmt.Printf("Input for day %d:\n%s\n", day, input)

	// Process the input as needed for Day 1
	input = strings.Trim(input, "\n ")
	fmt.Printf("Processed input for day %d:\n%s\n", day, input)
}
