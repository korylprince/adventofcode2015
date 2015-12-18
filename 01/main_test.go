package main

import "testing"

var findLevelTests = map[string]int{
	"(())":    0,
	"()()":    0,
	"(((":     3,
	"(()(()(": 3,
	")())())": -3,
}

func TestFindLevel(t *testing.T) {
	for test, val := range findLevelTests {
		if v := FindLevel(test); v != val {
			t.Error("Test \"%s\" Failed. Got %v, Expected %v", test, v, val)
		}
	}
}
