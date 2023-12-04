package solutions

import (
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

type ScratchCard struct {
    Id int
    Numbers []int
    WinningNumbers []int
}

type Node struct {
    Data *ScratchCard
    Next *Node
}

type Queue struct {
    Length int
    Head *Node
    Tail *Node
}

func (queue *Queue) Pop() *Node {
    node := queue.Head
    queue.Head = queue.Head.Next
    queue.Length -= 1
    return node
}

func (queue *Queue) Push(data *ScratchCard) {
    node := &Node{
        data,
        nil,
    }

    if queue.Head == nil {
        queue.Head = node
    }

    if queue.Tail != nil {
        queue.Tail.Next = node
    }
    queue.Tail = node
    queue.Length += 1
}

func (queue *Queue) Print() {
    node := queue.Head

    fmt.Print("{")
    for node != nil {
        fmt.Printf("#%d, ", node.Data.Id)
        node = node.Next
    }
    fmt.Println("}")
}

func Day4Solution() {
    fmt.Printf("----- Day 4 Solution -----\n")
    fmt.Println(" Part 1:", day4Part1())
    fmt.Println(" Part 2:", day4Part2())
    fmt.Printf("--------------------------\n")
}

func day4PuzzleInput() []*ScratchCard {
    cards := make([]*ScratchCard, 0)
    dat, err := os.ReadFile("./test/day4.txt")
    if err != nil {
        panic(err)
    }
    lines := strings.Split(string(dat), "\n")

    for _, line := range lines {
        if line == "" {
            continue
        }

        cardData := strings.Split(line, ": ")
        cardId := -1
        if val, err := strconv.Atoi(strings.Split(cardData[0], " ")[len(strings.Split(cardData[0], " "))-1]); err == nil {
            cardId = val
        }
        numbers := strings.Split(cardData[1], " | ")
        nums := make([]int, 0)
        winNums := make([]int, 0)

        for _, n := range strings.Split(strings.ReplaceAll(numbers[0], "  ", " "), " ") {
            if n == "" {
                continue
            }
            val, err := strconv.Atoi(n)
            if err != nil {
                panic(err)
            }
            nums = append(nums, val)
        }

        for _, n := range strings.Split(strings.ReplaceAll(numbers[1], "  ", " "), " ") {
            if n == "" {
                continue
            }

            val, err := strconv.Atoi(n)
            if err != nil {
                panic(err)
            }
            winNums = append(winNums, val)
        }

        cards = append(cards, &ScratchCard{
            cardId,
            nums,
            winNums,
        })
    }

    return cards
}

func day4Part1() int {
    sum := 0

    cards := day4PuzzleInput()
    for _, card := range cards {
        winMap := make(map[int]bool)
        matches := make([]int, 0)
        for _, num := range card.Numbers {
            winMap[num] = true
        }

        for _, winNum := range card.WinningNumbers {
            if winMap[winNum] {
                matches = append(matches, winNum)

            }
        }
        sum += int(math.Pow(2, float64(len(matches)-1)))
    }

    return sum
}

func day4Part2() int {
    sum := 0

    cards := day4PuzzleInput()
    matchesTable := make(map[int][]int)
    copiesTable := make(map[int]int)
    for _, card := range cards {
        winMap := make(map[int]bool)
        matches := make([]int, 0)
        for _, num := range card.Numbers {
            winMap[num] = true
        }

        for _, winNum := range card.WinningNumbers {
            if winMap[winNum] {
                matches = append(matches, winNum)

            }
        }

        matchesTable[card.Id] = matches
    }

    queue := &Queue{
        0,
        nil,
        nil,
    }

    for _, card := range cards {
        queue.Push(card)
        copiesTable[card.Id] += 1
    }

    for queue.Length > 0 {
        node := queue.Pop()

        for i := range matchesTable[node.Data.Id] {
            queue.Push(cards[node.Data.Id+i])
            copiesTable[node.Data.Id+i+1] += 1
        }
    }

    for _, v := range copiesTable {
        sum += v
    }

    return sum
}
