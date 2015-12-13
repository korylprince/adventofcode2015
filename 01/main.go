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

func main() {
	b, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(FindLevel(string(b)))
}
