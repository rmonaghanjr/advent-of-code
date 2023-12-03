package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GameInformation struct {
    Red int
    Blue int
    Green int
}

func max(a int, b int) int {
    if a > b {
        return a
    }

    return b
}

func min(a int, b int) int {
    if a < b {
        return a
    }

    return b
}

func Day2Solution() {
    fmt.Printf("----- Day 2 Solution -----\n")
    fmt.Println(" Part 1:", day2Part1())
    fmt.Println(" Part 2:", day2Part2())
    fmt.Printf("--------------------------\n")
}

func day2PuzzleInput() []string {
    dat, err := os.ReadFile("./test/day2.txt")
    if err != nil {
        panic(err)
    }
    lines := strings.Split(string(dat), "\n")

    return lines
}

func day2Part1() int {
    gameMap := make(map[int]*GameInformation)
    lines := day2PuzzleInput()

    for i, line := range lines {
        if !strings.Contains(line, "Game") {
            continue
        }
        gameMap[i+1] = &GameInformation {
            0,
            0, 
            0,
        }

        gameInfo := strings.Split(line, ": ")[1]
        games := strings.Split(gameInfo, "; ")
        for _, game := range games {
            cubesPulled := strings.Split(game, ", ")
            for _, cube := range cubesPulled {
                spl := strings.Split(cube, " ")
                n, err := strconv.Atoi(spl[0])
                if err != nil {
                    panic(err)
                }

                if spl[1] == "red" {
                    gameMap[i+1].Red = max(n, gameMap[i+1].Red)
                } else if spl[1] == "green" {
                    gameMap[i+1].Green = max(n, gameMap[i+1].Green)
                } else if spl[1] == "blue" {
                    gameMap[i+1].Blue = max(n, gameMap[i+1].Blue)
                }
            }
        }
    }

    sum := 0
    for id, info := range gameMap {
        if info.Red <= 12 && info.Green <= 13 && info.Blue <= 14 {
            sum += id
        }
    }

    return sum
}

func day2Part2() int {
    gameMap := make(map[int]*GameInformation)
    lines := day2PuzzleInput()

    for i, line := range lines {
        if !strings.Contains(line, "Game") {
            continue
        }
        gameMap[i+1] = &GameInformation {
            0,
            0,
            0,
        }

        gameInfo := strings.Split(line, ": ")[1]
        games := strings.Split(gameInfo, "; ")
        for _, game := range games {
            cubesPulled := strings.Split(game, ", ")
            for _, cube := range cubesPulled {
                spl := strings.Split(cube, " ")
                n, err := strconv.Atoi(spl[0])
                if err != nil {
                    panic(err)
                }

                if spl[1] == "red" {
                    gameMap[i+1].Red = max(n, gameMap[i+1].Red)
                } else if spl[1] == "green" {
                    gameMap[i+1].Green = max(n, gameMap[i+1].Green)
                } else if spl[1] == "blue" {
                    gameMap[i+1].Blue = max(n, gameMap[i+1].Blue)
                }
            }
        }
    }

    sum := 0
    for _, info := range gameMap {
        power := info.Red * info.Blue * info.Green
        sum += power
    }

    return sum
}

