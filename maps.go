package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	r := make(map[string]int)
	for _, p := range strings.Fields(s) {
		r[p]++
	}
	return r
}

func main() {
	wc.Test(WordCount)
}
