package solutions

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

type ScratchCard struct {
    Id int
    Matches int
}

type SCNode struct {
    Data *ScratchCard
    Next *SCNode
}

type SCStack struct {
    Length int
    Head *SCNode
}

func (stack *SCStack) Pop() *SCNode {
    node := stack.Head
    stack.Head = stack.Head.Next
    stack.Length -= 1
    return node
}

func (stack *SCStack) Push(data *ScratchCard) {
    node := &SCNode{
        data,
        stack.Head,
    }

    stack.Head = node
    stack.Length += 1
}

func (stack *SCStack) Print() {
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

    for i, line := range lines {
        if line == "" {
            continue
        }

        cardData := strings.Split(line, ": ")
        cardId := i+1
        numbers := strings.Split(cardData[1], " | ")
        nums := listAtoi(strings.Split(strings.ReplaceAll(numbers[0], "  ", " "), " "))
        winNums := listAtoi(strings.Split(strings.ReplaceAll(numbers[1], "  ", " "), " "))

        cards = append(cards, &ScratchCard{
            cardId,
            listIntersection(nums, winNums),
        })
    }

    return cards
}

func listIntersection(a []int, b []int) int {
    matches := 0
    lookup := make([]bool, 100)
    for _, av := range a {
        lookup[av] = true
    }

    for _, bv := range b {
        if !lookup[bv] { continue } 
        matches += 1
    }
    return matches
}

func day4Part1(cards []*ScratchCard) int {
    sum := 0
    for _, card := range cards {
        if card.Matches == 0 {
            sum += 1
            continue
        }
        sum += 1 << (card.Matches - 1)
    }
    return sum
}

func day4Part2(cards []*ScratchCard) int {
    sum := 0
    stack := &SCStack{0,nil}
    // build table of all matches for each card
    // add original copy of scratchers to the queue to be processed
    for _, card := range cards {
        stack.Push(card)
    }
    sum += stack.Length

    // bfs to visit all nodes in tree (there will be no cycles, so we can modify the algorithm)
    for stack.Length > 0 {
        node := stack.Pop()

        for i := 0; i < node.Data.Matches; i++ {
            stack.Push(cards[node.Data.Id+i])
        }
        sum += node.Data.Matches
    }

    return sum
}

