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
    cards := day4PuzzleInput()
    fmt.Printf("----- Day 4 Solution -----\n")
    fmt.Println(" Part 1:", day4Part1(cards))
    fmt.Println(" Part 2:", day4Part2(cards))
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

func listIntersection(a []int, b []int) []int {
    matches := make([]int, 0)
    lookup := make(map[int]bool)
    for _, av := range a {
        lookup[av] = true
    }

    for _, bv := range b {
        if lookup[bv] {
            matches = append(matches, bv)
        }
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

    queue := &Queue{
        0,
        nil,
        nil,
    }
    matchesTable := make(map[int][]int)
    // build table of all matches for each card
    // add original copy of scratchers to the queue to be processed
    for _, card := range cards {
        matchesTable[card.Id] = listIntersection(card.Numbers, card.WinningNumbers)
        queue.Push(card)
        sum += 1
    }

    // bfs to visit all nodes in tree (there will be no cycles, so we can modify the algorithm)
    for queue.Length > 0 {
        node := queue.Pop()

        for i := range matchesTable[node.Data.Id] {
            queue.Push(cards[node.Data.Id+i])
            sum += 1
        }
    }

    return sum
}
