package main

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/vinhpham2106/aoc25golang/utils"
	"math"
)

func main() {
	day := 2
	input, err := utils.GetInput(day)
	if err != nil {
		fmt.Printf("Error fetching input: %v\n", err)
		return
	}
	// DEBUG
	input = strings.Trim(input, "\n ")
	//input = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	fmt.Printf("Input for day %d:\n%s\n", day, input)

	ranges := strings.Split(input, ",")

	// Solve part 1
	ans1 := 0

	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		start := bounds[0]
		end := bounds[1]
		ans1 += part1(start, end)
	}

	fmt.Printf("Part 1: %d\n", ans1)
	// Solve part 2

	// Add some separation
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	ans2 := 0
	
	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		start := bounds[0]
		end := bounds[1]
		ans2 += part2(start, end)
	}

	ans2 += ans1
	
	fmt.Printf("Part 2: %d\n", ans2)
}

func part1(start string, end string) (invalid int) {
	maxInvalid := 0
	minInvalid := 0
	// Find min invalid id half

	if len(start) % 2 != 0 {
		minInvalid = int(math.Pow10(len(start) / 2))
	} else {
		halfLen := len(start) / 2
		leftPiece, _ := strconv.Atoi(start[:halfLen])
		rightPiece, _ := strconv.Atoi(start[halfLen:])
		if leftPiece >= rightPiece {
			minInvalid = leftPiece
		} else {
			minInvalid = leftPiece + 1
		}
	}

	// Find max invalid id

	if len(end) % 2 != 0 {
		maxInvalid = int(math.Pow10(len(end) / 2)) - 1
	} else {
		halfLen := len(end) / 2
		leftPiece, _ := strconv.Atoi(end[:halfLen])
		rightPiece, _ := strconv.Atoi(end[halfLen:])
		if leftPiece <= rightPiece {
			maxInvalid = leftPiece
		} else {
			maxInvalid = leftPiece - 1
		}
	}
	fmt.Printf("For range %s-%s, min invalid: %d, max invalid: %d\n", start, end, minInvalid, maxInvalid)
	ans := 0
	for i := minInvalid; i <= maxInvalid; i++ {
		ans += extendNum(i)
	}
	return ans
}

func part2(start string, end string) (invalid int) {
	numStart, _ := strconv.Atoi(start)
	numEnd, _ := strconv.Atoi(end)

	// Find min invalid id for extend3
	low := 1
	high := 9999
	minInvalid3 := 1
	maxInvalid3 := 0
	
	if numStart < 111 {
		minInvalid3 = 1
	} else {
		minInvalid3 = 10000
		for low <= high {
			mid := (low + high) / 2
			extended := extendNumOdd(mid, 3)
			if extended >= numStart {
				minInvalid3 = int(math.Min(float64(mid), float64(minInvalid3)))
				high = mid - 1
			} else {
				low = mid + 1
			}
		}

		if extendNumOdd(high, 3) >= numStart {
			minInvalid3 = int(math.Min(float64(high), float64(minInvalid3)))
		}
		if extendNumOdd(low, 3) >= numStart {
			minInvalid3 = int(math.Min(float64(low), float64(minInvalid3)))
		}
	}

	// Find max invalid id for extend3
	low = 1
	high = 9999

	if numEnd < 111 {
		maxInvalid3 = 0
	} else {
		for low <= high {
			mid := (low + high) / 2
			extended := extendNumOdd(mid, 3)
			if extended <= numEnd {
				maxInvalid3 = int(math.Max(float64(mid), float64(maxInvalid3)))
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		if extendNumOdd(high, 3) <= numEnd {
			maxInvalid3 = int(math.Max(float64(high), float64(maxInvalid3)))
		}
		if extendNumOdd(low, 3) <= numEnd {
			maxInvalid3 = int(math.Max(float64(low), float64(maxInvalid3)))
		}
	}

	fmt.Printf("For range %s-%s, min invalid 3: %d, max invalid 3: %d\n", start, end, minInvalid3, maxInvalid3)
	ans := 0
	for i := minInvalid3; i <= maxInvalid3; i++ {
		if !evenlySymmetric(i) {
			ans += extendNumOdd(i, 3)
			fmt.Printf("Adding invalid id %d extended to %d\n", i, extendNumOdd(i, 3))
		}
	}

	// Now do the same for extendOdd with mult 5

	low = 1
	high = 99

	minInvalid5 := 0
	maxInvalid5 := 0

	if numStart < 11111 {
		minInvalid5 = 1
	} else {
		minInvalid5 = 100
		for low <= high {
			mid := (low + high) / 2
			extended := extendNumOdd(mid, 5)
			if extended >= numStart {
				minInvalid5 = int(math.Min(float64(mid), float64(minInvalid5)))
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		if extendNumOdd(high, 5) >= numStart {
			minInvalid5 = int(math.Min(float64(high), float64(minInvalid5)))
		}
		if extendNumOdd(low, 5) >= numStart {
			minInvalid5 = int(math.Min(float64(low), float64(minInvalid5)))
		}
	}

	// Find max invalid id for extend5
	low = 1
	high = 99

	if numEnd < 11111 {
		maxInvalid5 = 0
	} else {
		for low <= high {
			mid := (low + high) / 2
			extended := extendNumOdd(mid, 5)
			if extended <= numEnd {
				maxInvalid5 = int(math.Max(float64(mid), float64(maxInvalid5)))
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		if extendNumOdd(high, 5) <= numEnd {
			maxInvalid5 = int(math.Max(float64(high), float64(maxInvalid5)))
		}
		if extendNumOdd(low, 5) <= numEnd {
			maxInvalid5 = int(math.Max(float64(low), float64(maxInvalid5)))
		}
	}

	fmt.Printf("For range %s-%s, min invalid 5: %d, max invalid 5: %d\n", start, end, minInvalid5, maxInvalid5)
	for i := minInvalid5; i <= maxInvalid5; i++ {
		if !evenlySymmetric(i) {
			ans += extendNumOdd(i, 5)
			fmt.Printf("Adding invalid id %d extended to %d\n", i, extendNumOdd(i, 5))
		}
	}

	// Define int array of all possible invalid mutl 7
	theSeven7 := [9]int{1111111, 2222222, 3333333, 4444444, 5555555, 6666666, 7777777, 8888888, 9999999}
	if numStart >= 1111111 {
		for _, val := range theSeven7 {
			if val >= numStart && val <= numEnd {
				ans += val
				fmt.Printf("Adding invalid id %d\n", val)
			}
		}
	}
	return ans
}

func extendNum(num int) int {
	length := int(math.Log10(float64(num))) + 1
	ans := int(math.Pow10(length)) * num + num
	//fmt.Printf("Extend %d to %d\n", num, ans)
	return ans
}

func extendNumOdd(num int, mult int) int {
	length := int(math.Log10(float64(num))) + 1
	newLength := length
	ans := num

	for i := 1; i < mult; i++ {
		ans = int(math.Pow10(length)) * ans + num
		newLength += length
	}
	return ans
}

func evenlySymmetric(num int) bool {
	length := int(math.Log10(float64(num))) + 1
	if length % 2 != 0 {
		return false
	}
	halfLength := (int(math.Log10(float64(num))) + 1) / 2
	powTen := int(math.Pow10(halfLength))
	left := num / powTen
	right := num % powTen
	return left == right
}
// 122792346
// 1227775554