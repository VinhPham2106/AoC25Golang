package main

import (
	"fmt"

	"github.com/vinhpham2106/aoc25golang/utils"
)

func main() {
	day := 1 // Example day
	input, err := utils.GetInput(day)
	if err != nil {
		fmt.Printf("Error fetching input: %v\n", err)
		return
	}
	fmt.Printf("Input for day %d:\n%s\n", day, input)
}
