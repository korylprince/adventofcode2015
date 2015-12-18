package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func FindLevel(in string) int {
	return strings.Count(in, "(") - strings.Count(in, ")")
}

func FindLevelIndex(in string, level int) int {
	var cur int
	for idx, val := range in {
		if val == '(' {
			cur += 1
		} else if val == ')' {
			cur -= 1
		} else {
			panic(fmt.Errorf("Unexpected '%s'", val))
		}
		if cur == level {
			return idx + 1 //starts at 1
		}
	}
	panic(fmt.Errorf("Reached EOF without finding index"))
}

func main() {
	b, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answer 1:", FindLevel(string(b)))
	fmt.Println("Answer 2:", FindLevelIndex(string(b), -1))
}
