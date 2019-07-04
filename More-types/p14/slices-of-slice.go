package p14

import (
	"fmt"
	"strings"
)

func Main() {
	board := [][]string {
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}
	board[0][0] = "X"  // 0 row 0 column = X
	board[0][2] = "O"  // 0 row 2 column = O
	// board[1][0] = "X"
	// board[0][1] = "O"
	// board[0][0] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}