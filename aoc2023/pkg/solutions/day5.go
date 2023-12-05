package solutions

import (
	"fmt"
	"math"
	"os"
	"strings"
	"sync"
)

type Almanac struct {
    Seeds []int
    LookupMap map[string]*RangedMap
}

type RangedMap struct {
    From string
    To string
    Ranges []*Range
}

type Range struct {
    DestinationStart int
    SourceStart int
    Length int
}

var APPLICATION_ORDER = []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}

func Day5Solution() {
    fmt.Printf("----- Day 5 Solution -----\n")
    fmt.Println(" Part 1:", day5Part1())
    fmt.Println(" Part 2:", day5Part2())
    fmt.Printf("--------------------------\n")
}

func day5PuzzleInput() *Almanac {
    almanac := &Almanac{
        make([]int, 0),
        make(map[string]*RangedMap),
    }
    dat, err := os.ReadFile("./test/day5.txt")
    if err != nil {
        panic(err)
    }
    lines := strings.Split(string(dat), "\n\n")

    for _, m := range lines {
        if m == "" { continue }
        if strings.Contains(m, "seeds") {
            almanac.Seeds = listAtoi(strings.Split(strings.Split(m, ": ")[1], " "))
            continue
        }

        sLines := strings.Split(m, "\n")
        key := strings.ReplaceAll(sLines[0], ":", "")
        key = strings.ReplaceAll(key, " map", "")
        key = strings.ReplaceAll(key, "to-", "")
        broken := strings.Split(key, "-")
        from := broken[0]
        to := broken[1]
        almanac.LookupMap[to] = &RangedMap {
            from,
            to,
            make([]*Range, 0),
        }
        for j, rm := range sLines {
            if rm == "" { continue }
            if j == 0 { continue }
            ns := listAtoi(strings.Split(rm, " "))
            rng := &Range{
                ns[0],
                ns[1],
                ns[2],
            }
            almanac.LookupMap[to].Ranges = append(almanac.LookupMap[to].Ranges, rng)
        }
    }

    return almanac 
}

func listMin(nums []int) int {
    minimum := math.MaxInt
    for _, n := range nums {
        if n < minimum { minimum = n }
    }
    return minimum
}

func getNumberInNextMap(n int, rangedMap *RangedMap) int {
    for _, rng := range rangedMap.Ranges {
        if n >= rng.SourceStart && n < rng.SourceStart + rng.Length {
            return rng.DestinationStart + (n - rng.SourceStart)
        }
    }

    return n
}

func fillArray(start int, count int) []int {
    arr := make([]int, count)
    for i := 0; i < count; i++ {
        arr[i] = start + i
    }
    return arr
}

func day5Part1() int {
    input := day5PuzzleInput()
    location := math.MaxInt
    for _, seed := range input.Seeds {
        foundN := seed
        for _, v := range APPLICATION_ORDER {
            m := input.LookupMap[v]
            foundN = getNumberInNextMap(foundN, m)
        } 
        
        if foundN < location { location = foundN }
    }

    return location
}

func day5Part2() int {
    input := day5PuzzleInput()
    final := math.MaxInt
    channel := make(chan int)
    wg := &sync.WaitGroup{}
    for i := 0; i < len(input.Seeds); i+=2 {
        wg.Add(1)
        go func(seed int, length int, wg *sync.WaitGroup, c chan int) {
            defer wg.Done()
            location := math.MaxInt
            arr := fillArray(seed, length)
            for i := 0; i < len(arr); i++ {
                foundN := arr[i]
                for _, v := range APPLICATION_ORDER {
                    m := input.LookupMap[v]
                    foundN = getNumberInNextMap(foundN, m)
                } 

                if foundN < location { location = foundN }
            }
            c <- location
        }(input.Seeds[i], input.Seeds[i+1], wg, channel)
    }

    go func(wg *sync.WaitGroup, c chan int) {
        wg.Wait()
        close(c)
    }(wg, channel)

    for v := range channel {
        if v < final { final = v }
    }

    return final
}


