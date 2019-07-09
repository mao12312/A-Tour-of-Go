package p23

import (
	"golang.org/x/tour/wc"
)

func WordCount (s string) map[string]int {
	return map[string]int{"x": 1}
}

func Main() {
	wc.Test(WordCount)
}