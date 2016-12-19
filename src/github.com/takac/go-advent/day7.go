package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func IsAbba(s string) bool {
	// fmt.Println("check:", s)
	for i := 1; i < len(s)-2; i++ {
		if s[i-1] == s[i+2] && s[i] == s[i+1] && s[i] != s[i-1] {
			// fmt.Printf("%s\n", s[i-1:i+3])
			return true
		}
	}
	return false
}

func GetHyper(s string) []string {
	for {
		idxOpen := strings.Index(s, "[")
		if idxOpen > -1 {
			idxClose := strings.Index(s, "]")
			h := s[idxOpen+1 : idxClose]
			return append(GetHyper(s[idxClose+1:]), h)
		}
		return make([]string, 0)
	}
}

func CheckHyper(s string) bool {
	for {
		idxOpen := strings.Index(s, "[")
		if idxOpen > -1 {
			idxClose := strings.Index(s, "]")
			hyper := IsAbba(s[idxOpen+1 : idxClose])
			if hyper {
				return false
			}
			return CheckHyper(s[idxClose+1:])
		}
		return true
	}
}

func GetSupernet(s string) []string {
	for {
		idxOpen := strings.Index(s, "[")
		if idxOpen > -1 {
			idxClose := strings.Index(s, "]")
			h := s[:idxOpen]
			return append(GetSupernet(s[idxClose+1:]), h)
		}
		return []string{s}
	}
}

func CheckSupernet(s string) bool {
	for {
		idxOpen := strings.Index(s, "[")
		if idxOpen > -1 {
			idxClose := strings.Index(s, "]")
			return CheckSupernet(s[idxClose+1:]) || IsAbba(s[:idxOpen])
		}
		return IsAbba(s)
	}
}

func FindBABs(nets []string) []string {
	babs := make([]string, 0)
	for _, net := range nets {
		for i := 1; i < len(net)-1; i++ {
			if net[i-1] == net[i+1] && net[i] != net[i-1] {
				babs = append(babs, net[i-1:i+2])
			}
		}
	}
	return babs
}

func BABwithABA(line string) bool {
	babs := FindBABs(GetHyper(line))
	abas := FindBABs(GetSupernet(line))
	for _, bab := range babs {
		for _, aba := range abas {
			if aba[0] == bab[1] && aba[1] == bab[0] {
				fmt.Println(bab, aba)
				return true
			}
		}
	}
	return false
}

func Abba(s string) bool {
	return CheckHyper(s) && CheckSupernet(s)
}

func main() {
	// fmt.Println(Abba("abba[mnop]qrst"))
	// fmt.Println(Abba("abcd[bddb]xyyx"))
	// fmt.Println(Abba("aaaa[qwer]tyui"))
	fmt.Println(FindBABs(GetHyper("ioxxoj[asdfgh]zxcvbn[xaax]accd")))
	fmt.Println(FindBABs(GetSupernet("ioxxoj[asdfgh]zxcvbn[xaax]accd")))
	// inc := 0
	// dat, err := ioutil.ReadFile("advent7.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
	// 	if Abba(strings.TrimSpace(line)) {
	// 		fmt.Println("Valid:", line)
	// 		inc++
	// 	}
	// }
	// fmt.Println(inc)
	inc := 0
	dat, err := ioutil.ReadFile("advent7.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		if BABwithABA(strings.TrimSpace(line)) {
			fmt.Println("Valid:", line)
			inc++
		}
	}
	fmt.Println(inc)
}
