package main

import (
    "log"
    "fmt"
    "strings"
    "io/ioutil"
)

type Point struct {
    x, y int
}

type Direction int

func (e Direction) Point() Point {
	switch e {
	case Up:
		return Point{-1, 0}
	case Right:
        return Point{0, 1}
	case Down:
        return Point{1, 0}
	case Left:
        return Point{0, -1}
	}
    log.Fatal("No valid string representation")
    return Point{}
}

func (e Direction) String() string {
	switch e {
	case Up:
		return "Up"
	case Right:
		return "Right"
	case Down:
		return "Down"
	case Left:
		return "Left"
	}
    log.Fatal("No valid string representation")
    return ""
}

const (
    Up Direction = iota
    Right
    Down
    Left
)

func (o Point) addDirection(dir Direction, mask [][]string) Point {
    var p Point
	switch dir {
	case Up:
		p = Point{-1, 0}
	case Right:
        p = Point{0, 1}
	case Down:
        p = Point{1, 0}
	case Left:
        p = Point{0, -1}
    default:
        log.Fatal("No valid representation")
        return Point{}
	}
    combine := p.add(o, len(mask[0])-1)
    if mask[combine.x][combine.y] == "_" {
        return o
    }
    return combine
}


func (o Point) add(p Point, bounds int) Point {
    tx := p.x + o.x
    ty := p.y + o.y
    if tx > bounds {
        tx = bounds
    }
    if tx < 0 {
        tx = 0
    }
    if ty > bounds {
        ty = bounds
    }
    if ty < 0 {
        ty = 0
    }
    return Point{tx,ty}
}

func CharToDirection(c rune) Direction {
    switch c {
    case 'U': return Up
    case 'D': return Down
    case 'L': return Left
    case 'R': return Right
    default:
        log.Fatal("No valid Direction representation:", c)
        return -1
    }
}

func (p Point) toNum(size int) int {
    return (size * p.x) + p.y + 1
}

func (p Point) fromMask(mask [][]string) string {
    return mask[p.x][p.y]
}

func __main() {
    // mask := make([][]string, 3)
    // mask[0] = []string{"1","2","3"}
    // mask[1] = []string{"4","5","6"}
    // mask[2] = []string{"7","8","9"}
    mask := make([][]string, 5)
    mask[0] = []string{"_","_","1","_","_"}
    mask[1] = []string{"_","2","3","4","_"}
    mask[2] = []string{"5","6","7","8","9"}
    mask[3] = []string{"_","A","B","C","_"}
    mask[4] = []string{"_","_","D","_","_"}
    // max array bounds
    // bounds := 2
    // 3x + y + 1
    start := Point{2,0}

    dat, err := ioutil.ReadFile("advent2.txt")
    if err != nil {
        panic(err)
    }
    s := strings.Split(strings.TrimSpace(string(dat)), "\n")
    // fmt.Println(start, start.toNum(3))
    for _, line := range s {
        // fmt.Println("Line", i, line)
        for _, c := range line {
            dir := CharToDirection(c)
            start = start.addDirection(dir, mask)
            // fmt.Println(start, start.fromMask(mask))
        }
        fmt.Println(start.fromMask(mask))
    }
}
