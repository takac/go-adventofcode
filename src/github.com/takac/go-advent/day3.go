package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

func IsTriangle(side1, side2, side3 int) bool {
     if (side1 + side2) > side3 && (side1 + side3) > side2 && (side2 + side3) > side1 {
         fmt.Println("valid:", side1, side2, side3)
         return true
     }
     fmt.Println("invalid:", side1, side2, side3)
     return false
 }


func ___main() {
    dat, err := ioutil.ReadFile("advent3.txt")
    if err != nil {
        panic(err)
    }

    doc := strings.Split(strings.TrimSpace(string(dat)), "\n")
    counter := 0
    db := make([][]int, 1908)
    for i, line := range doc {
         tri := strings.Split(line, " ")
         row1,_ := strconv.Atoi(tri[0])
         row2,_ := strconv.Atoi(tri[1])
         row3,_ := strconv.Atoi(tri[2])
         db[i] = []int{row1,row2,row3}
         if i % 3 == 2 {
             for j:=0; j<3; j++ {
                 fmt.Println("col", db[i-2][j], db[i-1][j], db[i][j])
                 if IsTriangle(db[i-2][j], db[i-1][j], db[i][j]) {
                     counter++
                 }
             }
         }

    }

    fmt.Println("total valid:", counter)

}
