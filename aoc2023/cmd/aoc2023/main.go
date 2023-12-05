package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rmonaghanjr/advent-of-code/aoc2023/pkg/solutions"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: ", os.Args[0], " <days>")
        os.Exit(1)
    }

    days := strings.Split(os.Args[1], ",")
    for i, d := range days {
        switch d {
        case "day1": solutions.Day1Solution()
        case "day2": solutions.Day2Solution()
        case "day3": solutions.Day3Solution()
        case "day4": solutions.Day4Solution()
        case "day5": solutions.Day5Solution()
        default:
            fmt.Printf("Day #%d has not been implemented!\n", i+1)
        }
    }
}
