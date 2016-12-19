package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func old_main() {
	start := "ffykfhsq"
	password := []byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	for i := 0; ; i++ {
		data := []byte(start + strconv.Itoa(i))
		out := fmt.Sprintf("%x", md5.Sum(data))
		if out[0:5] == "00000" {
			if out[5] > 47 && out[5] < 56 {
				fmt.Printf("%c - %s - %d\n", out[5], out, i)
				n, _ := strconv.Atoi(string(out[5]))
				if password[n] == ' ' {
					password[n] = out[6]
				}
			}
		}
		complete := true
		for _, c := range password {
			if c == ' ' {
				complete = false
				break
			}
		}
		if complete {
			fmt.Println(string(password[:]))
			break
		}
	}
}
