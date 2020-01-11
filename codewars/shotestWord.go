package shortestWord

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"sort"
)

func FindShort(s string)int {
	var stringCount []int
	slice := strings.Split(s, " ")
	for _, s := range slice {
		stringCount = append(stringCount, utf8.RuneCountInString(s))
	}
	sort.Slice(stringCount, func(i, j int) bool {
		return stringCount[i] < stringCount[j]
		})
	return stringCount[0]
	
}


// package kata

// import "strings"

// func FindShort(s string) int {
//   shortest := len(s)
//   for _, word := range strings.Split(s, " ") {
//     if len(word) < shortest {
//       shortest = len(word)
//     }
//   }
//   return shortest
// }
