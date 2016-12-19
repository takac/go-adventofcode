package main

import (
    "strings"
    "strconv"
    "container/heap"
    "io/ioutil"
    "fmt"
)
type LCHeap []LetterCount

func (h LCHeap) Len() int           { return len(h) }

func (h LCHeap) Less(i, j int) bool {
    if h[i].Count < h[j].Count {
        return false
    }
    if h[i].Count > h[j].Count {
        return true
    }
    if h[i].Count == h[j].Count {
        if h[i].Letter > h[j].Letter {
            return false
        }
    }
    return true
}

func (h LCHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *LCHeap) Push(x interface{}) {
	*h = append(*h, x.(LetterCount))
}

func (h *LCHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type LetterCount struct {
    Letter rune
    Count int
}

func (lc LetterCount) String() string {
    return fmt.Sprintf("%c:%d", lc.Letter, lc.Count)
}

func main() {

    dat, err := ioutil.ReadFile("advent4.txt")
    if err != nil {
        panic(err)
    }

    doc := strings.Split(strings.TrimSpace(string(dat)), "\n")
    fmt.Println("len", len(doc))
    numberSum := 0
    for _, line := range doc {
        counter := make(map[rune]*LetterCount)
        checksum := line[len(line)-6:len(line)-1]
        number := line[len(line)-10:len(line)-7]
        for _, letter := range line[0:len(line)-11] {
            lc, ok := counter[letter]
            if ! ok {
                // fmt.Printf("New counter for %c\n", letter)
                lc = &LetterCount{letter, 0}
                counter[letter] = lc
            }
            lc.Count = lc.Count + 1
            // fmt.Printf("%s\n", lc)
        }
        delete(counter, '-')
        arrayLen := len(counter)
        myHeap := make(LCHeap, arrayLen)
        inc := 0
        for _, value := range counter {
            myHeap[inc] = *value
            inc++
        }
        heap.Init(&myHeap)
        var chck string
        for i:=0; i<5; i++ {
            chck = chck + string(heap.Pop(&myHeap).(LetterCount).Letter)
        }
        if chck == checksum {
            // fmt.Println(number)
            iNumber, _ := strconv.Atoi(number)
            numberSum = numberSum + iNumber
            // fmt.Printf("%s\n", checksum)
            // fmt.Printf("%s\n", chck)
        }
        // fmt.Println(myHeap)
        // fmt.Println()
    }
    fmt.Println(numberSum)
}


