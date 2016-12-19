package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func _X_main() {
	maparray := make([]map[rune]int, 8)
	for i := 0; i < 8; i++ {
		maparray[i] = make(map[rune]int)
	}
	dat, err := ioutil.ReadFile("advent6.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		line = strings.TrimSpace(line)
		for i, c := range line {
			count, ok := maparray[i][c]
			if !ok {
				maparray[i][c] = 1
			} else {
				maparray[i][c] = count + 1
			}
			// fmt.Printf("%c\n", c)
		}
		// fmt.Println(line)
	}

	for _, m := range maparray {
		top := math.MaxInt64
		var mode rune
		for k, v := range m {
			if v < top {
				top = v
				mode = k
			}
		}
		fmt.Printf("%c", mode)
	}

}
