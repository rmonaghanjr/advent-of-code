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

type Stack struct {
    Length int
    Head *Node
}

func InitStack() *Stack {
    return &Stack{0,nil}
}

func (stack *Stack) Pop() *Node {
    node := stack.Head
    stack.Head = stack.Head.Next
    stack.Length -= 1
    return node
}

func (stack *Stack) Push(data *ScratchCard) {
    node := &Node{
        data,
        stack.Head,
    }

    stack.Head = node
    stack.Length += 1
}

func (stack *Stack) Print() {
    node := stack.Head

    fmt.Print("{")
    for node != nil {
        fmt.Printf("#%d, ", node.Data.Id)
        node = node.Next
    }
    fmt.Println("}")
}

func Day4Solution() {
    cards := day4PuzzleInput()
    fmt.Printf("----- Day 4 Solution -----\n")
    fmt.Println(" Part 1:", day4Part1(cards))
    fmt.Println(" Part 2:", day4Part2(cards))
    fmt.Printf("--------------------------\n")
}

func listAtoi(nums []string) []int {
    converted := make([]int, 0)
    for _, n := range nums {
        if n == "" { continue }

        val, err := strconv.Atoi(n)
        if err != nil {
            panic(err)
        }
        converted = append(converted, val)
    }
    return converted
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
        cardIdent := strings.Split(cardData[0], " ")
        cardId := -1
        if val, err := strconv.Atoi(cardIdent[len(cardIdent)-1]); err == nil {
            cardId = val
        }
        numbers := strings.Split(cardData[1], " | ")
        nums := listAtoi(strings.Split(strings.ReplaceAll(numbers[0], "  ", " "), " "))
        winNums := listAtoi(strings.Split(strings.ReplaceAll(numbers[1], "  ", " "), " "))

        cards = append(cards, &ScratchCard{
            cardId,
            nums,
            winNums,
        })
    }

    return cards
}

func listIntersection(a []int, b []int) []int {
    matches := make([]int, 0)
    lookup := make(map[int]bool)
    for _, av := range a {
        lookup[av] = true
    }

    for _, bv := range b {
        if !lookup[bv] { continue } 
        matches = append(matches, bv)
    }
    return matches
}

func day4Part1(cards []*ScratchCard) int {
    sum := 0

    for _, card := range cards {
        matches := listIntersection(card.Numbers, card.WinningNumbers)
        sum += int(math.Pow(2, float64(len(matches)-1)))
    }

    return sum
}

func day4Part2(cards []*ScratchCard) int {
    sum := 0

    stack := InitStack()
    // build table of all matches for each card
    // add original copy of scratchers to the queue to be processed
    matchesTable := make(map[int][]int)
    for _, card := range cards {
        matchesTable[card.Id] = listIntersection(card.Numbers, card.WinningNumbers)
        stack.Push(card)
        sum += 1
    }

    // bfs to visit all nodes in tree (there will be no cycles, so we can modify the algorithm)
    for stack.Length > 0 {
        node := stack.Pop()

        for i := range matchesTable[node.Data.Id] {
            stack.Push(cards[node.Data.Id+i])
            sum += 1
        }
    }

    return sum
}

