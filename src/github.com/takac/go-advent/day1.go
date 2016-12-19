package main

import (
    "log"
    // "container/list"
    "fmt"
    "strconv"
    "io/ioutil"
    "strings"
)

type Dir int

func (e Dir) String() string {
	switch e {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	}
    log.Fatal("No valid string representation")
    return ""
}

const (
    North Dir = iota
    East
    South
    West
)

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func _main() {
    dat, err := ioutil.ReadFile("advent.txt")
    if err != nil {
        panic(err)
    }

    s := strings.Split(string(dat), ",")
    direction := North
    counter := [4]int{0,0,0,0}
    // visited := list.New()
    // visited.PushFront([2]int{0,0})
    visitedMap := make(map[[2]int]bool)
    visitedMap[[2]int{0,0}] = true
    for _, move := range s {
        move = strings.TrimSpace(move)
        switch move[0] {
        case 'R':
            n, _ := strconv.Atoi(move[1:])
            direction = (direction + 1) % 4
            fmt.Println("R", n, direction)

            for i:=0; i < n; i++ {
                counter[direction] = counter[direction] + 1
                horz := counter[1] - counter[3]
                vert := counter[0] - counter[2]
                point := [2]int{horz, vert}
                fmt.Println("point:", point)
                if _, ok := visitedMap[point]; ok {
                    fmt.Println("VISITED BEFORE:", point)
                }
                visitedMap[point] = true
            }
        case 'L':
            n, _ := strconv.Atoi(move[1:])
            direction = (direction + 3) % 4
            fmt.Println("L", n, direction)
            for i:=0; i < n; i++ {
                counter[direction] = counter[direction] + 1
                horz := counter[1] - counter[3]
                vert := counter[0] - counter[2]
                point := [2]int{horz, vert}
                fmt.Println("point:", point)
                if _, ok := visitedMap[point]; ok {
                    fmt.Println("VISITED BEFORE:", point)
                }
                visitedMap[point] = true
            }
        default:
            log.Fatal("problem in parser for", move, "thing")
            panic("parser issue")
        }
        // visited.PushFront(point)
    }

    for i:=North; i < 4; i++ {
        fmt.Println(i, counter[i])
    }

    // for e := visited.Back(); e != nil; e = e.Prev() {
    //    if i, ok := e.Value.([2]int); ok {
    //        fmt.Println(i)
    //    }
   // }
    vert := abs(counter[0] - counter[2])
    horz := abs(counter[1] - counter[3])
    fmt.Println("vertical distance:", vert)
    fmt.Println("horizontal distance:", horz)
    fmt.Println("total distance:", horz+vert)
}
