package solutions

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Day1Solution() {
	fmt.Printf("----- Day 1 Solution -----\n")
    fmt.Println(" Part 1:", day1Part1())
    fmt.Println(" Part 2:", day2Part2())
    fmt.Printf("--------------------------\n")
}

func day1PuzzleInput() []string {
    dat, err := os.ReadFile("./test/day1_test.txt")
    if err != nil {
        panic(err)
    }
    lines := strings.Split(string(dat), "\n")

    return lines
}

func day1Part1() int {
    nums := make([]int, len(day1PuzzleInput()))
    sum := 0
    for _, line := range day1PuzzleInput() {
        if line == "" {
            break 
        }
        first := -1 
        last := -1
        for _, c := range []byte(line) {
            if int(c) - int('0') < 10 {
                if first == -1 {
                    first = int(c) - int('0')
                    first *= 10
                }   
                last = int(c) - int('0')
            } 
        }
        nums = append(nums,  first + last)
    }
    
    for _, num := range nums {
        sum += num
    }

    return sum
}

func toNumber(dgt string) int {
    digit := -1
    switch dgt {
    case "one": digit = 1
    case "two": digit = 2
    case "three": digit = 3
    case "four": digit = 4
    case "five": digit = 5
    case "six": digit = 6
    case "seven": digit = 7
    case "eight": digit = 8
    case "nine": digit = 9
    }

    if digit == -1 {
        digit = int(([]byte(dgt))[0]) - int('0')
    }

    return digit
}

func day1Part2() int {
    input := day1PuzzleInput()
    nums := make([]int, len(input))
    sum := 0
    for _, line := range input {
        r := regexp.MustCompile(`(?:(?=\d+)(?:zero|one|two|three|four|five|six|seven|eight|nine)|\d+(\.\d+)?|\w+)`)
        matches := r.FindAllStringSubmatch(line, -1)

        first := toNumber(matches[0][1]) * 10
        last := toNumber(matches[len(matches) - 1][1])

        nums = append(nums, first + last)
    }

    for _, num := range nums {
        sum += num
    }

    return sum
}
