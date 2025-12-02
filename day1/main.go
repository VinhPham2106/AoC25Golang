package main

import (
	"fmt"
	"strings"
	"strconv"
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
	//input = "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82"
	inputs := strings.Split(input, "\n")
	//fmt.Printf("%v\n", inputs)

	// Solve part 1

	curr := 50 
	dial := 100

	ans1 := 0

	for _, rot := range inputs {
		op := rot[0]
		count, err := strconv.Atoi(rot[1:])
		if err != nil {
			fmt.Printf("Error parsing count: %v\n", err)
			return
		}
		count %= dial
		if op == 'R' {
			curr += count
		} else if op == 'L' {
			curr += dial - count
		}
		curr %= dial
		if curr == 0 {
			ans1 += 1
		}
	}
	fmt.Printf("Part 1: %d\n", ans1)

	// Solve part 2

	ans2 := 0
	curr = 50
	
	for _, rot := range inputs {
		op := rot[0]
		count, err := strconv.Atoi(rot[1:])
		if err != nil {
			fmt.Printf("Error parsing count: %v\n", err)
			return
		}
		if op == 'R' {
			toZero := dial - curr
			if count < toZero {
				curr += count
			} else {
				count -= toZero
				ans2 += 1
				curr = 0
				ans2 += count / 100
				curr = count % 100
			}
		} else if op == 'L' {
			toZero := curr
			if count < toZero {
				curr -= count
			} else {
				count -= toZero
				if curr != 0 {
					ans2 += 1
				}
				curr = 0
				ans2 += count / 100
				curr = 100 - count % 100
			}
		}
		curr %= dial
		// fmt.Printf("Current dial position: %d\n", curr)
		// fmt.Printf("Current ans2 value: %d\n", ans2)
	}

	fmt.Printf("Part 2: %d\n", ans2)
}
