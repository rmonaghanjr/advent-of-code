package solutions

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func Day3Solution() {
    fmt.Printf("----- Day 3 Solution -----\n")
    fmt.Println(" Part 1:", day3Part1())
    fmt.Println(" Part 2:", day3Part2())
    fmt.Printf("--------------------------\n")
}

func day3PuzzleInput() []string {
    dat, err := os.ReadFile("./test/day3.txt")
    if err != nil {
        panic(err)
    }
    lines := strings.Split(string(dat), "\n")

    return lines
}

type Coord struct {
    Y int
    X int
    Width int
    Value string 
}

type Pair struct {
    RatioA int
    RatioB int
}

func numberIsNextToSymbol(x int, y int, width int, grid []string) bool {
    coordsToCheck := make([]*Coord, 0)
    coordsToCheck = append(coordsToCheck, &Coord{-1,0,0,""})
    coordsToCheck = append(coordsToCheck, &Coord{-1,-1,0,""})
    coordsToCheck = append(coordsToCheck, &Coord{-1,1,0,""})
    coordsToCheck = append(coordsToCheck, &Coord{width,0,0,""})
    coordsToCheck = append(coordsToCheck, &Coord{width,-1,0,""})
    coordsToCheck = append(coordsToCheck, &Coord{width,1,0,""})
    for i := 0; i < width; i++ {
        coordsToCheck = append(coordsToCheck, &Coord{i, 1,0,""})
        coordsToCheck = append(coordsToCheck, &Coord{i, -1,0,""})
    }

    for _, coord := range coordsToCheck {
        if grid[x+coord.X][y+coord.Y] != '.' {
            return true
        }
    }

    return false 
}

func numberIsNextToStar(x int, y int, width int, grid []string) *Coord {
    coordsToCheck := make([]*Coord, 0)
    coordsToCheck = append(coordsToCheck, &Coord{-1,0,0,""})
    coordsToCheck = append(coordsToCheck, &Coord{-1,-1,0,""})
    coordsToCheck = append(coordsToCheck, &Coord{-1,1,0,""})
    coordsToCheck = append(coordsToCheck, &Coord{width,0,0,""})
    coordsToCheck = append(coordsToCheck, &Coord{width,-1,0,""})
    coordsToCheck = append(coordsToCheck, &Coord{width,1,0,""})
    for i := 0; i < width; i++ {
        coordsToCheck = append(coordsToCheck, &Coord{i, 1,0,""})
        coordsToCheck = append(coordsToCheck, &Coord{i, -1,0,""})
    }

    for _, coord := range coordsToCheck {
        if grid[x+coord.X][y+coord.Y] == '*' {
            return &Coord {x+coord.X, y+coord.Y, 0, "*"} 
        }
    }

    return nil
}


func findPositionsAndWidthOfNumbers(grid []string) []*Coord {
    positions := make([]*Coord, 0)
    for i, line := range grid {
        if line == "" {
            continue
        }
        isInNumber := false
        for j, char := range line {
            if !isInNumber && char >= '0' && char <= '9' {
                isInNumber = true
                positions = append(positions, &Coord{j, i, 1, string(char)})
                continue
            }

            if isInNumber {
                if char >= '0' && char <= '9' {
                    positions[len(positions)-1].Width += 1
                    positions[len(positions)-1].Value += string(char)
                } else {
                    isInNumber = false
                }
            }
        }
    }

    return positions
}

func day3Part1() int {
    sum := 0
    input := day3PuzzleInput()
    grid := make([]string, len(input)+2)
    grid = append(grid, strings.Repeat(".", len(input[0])+2))
    for _, value := range input {
        if value == "" {
            continue
        }
        grid = append(grid, "." + value + ".")
    }
    grid = append(grid, strings.Repeat(".", len(input[0])+2))

    for _, number := range findPositionsAndWidthOfNumbers(grid) {
        isValidPart := numberIsNextToSymbol(number.X, number.Y, number.Width, grid)
        if isValidPart {
            if val, err := strconv.Atoi(number.Value); err == nil {
                sum += val
            }
        }
    }

    return sum
}

func day3Part2() int {
    sum := 0
    gearsMap := make(map[string]*Pair, 0)
    input := day3PuzzleInput()
    grid := make([]string, len(input)+2)
    grid = append(grid, strings.Repeat(".", len(input[0])+2))
    for _, value := range input {
        if value == "" {
            continue
        }
        grid = append(grid, "." + value + ".")
    }
    grid = append(grid, strings.Repeat(".", len(input[0])+2))

    for _, number := range findPositionsAndWidthOfNumbers(grid) {
        starCoord := numberIsNextToStar(number.X, number.Y, number.Width, grid)
        if starCoord != nil {
            key := strconv.Itoa(starCoord.X)+","+strconv.Itoa(starCoord.Y)
            if gearsMap[key] == nil {
                gearsMap[key] = &Pair {-1,-1} 
            }

            if val, err := strconv.Atoi(number.Value); err == nil {
                if gearsMap[key].RatioA == -1 {
                    gearsMap[key].RatioA = val 
                } else if gearsMap[key].RatioB == -1 {
                    gearsMap[key].RatioB = val 
                } else {
                    gearsMap[key].RatioA = 0
                    gearsMap[key].RatioB = 0
                }
            }
        }
    }

    for _, v := range gearsMap {
        if v.RatioA < 0 || v.RatioB < 0 {
            continue
        } 
        sum += v.RatioA * v.RatioB
    }

    return sum
}


