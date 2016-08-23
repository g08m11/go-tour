// goのmoretypesの学びの内容は全てこのファイルで行う
package main

import (
  "fmt"
  "strings"
)

func main() {

  board := [][]string {
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
  }

  board[0][0] = "X"
  board[2][2] = "o"
  board[1][2] = "X"
  board[1][0] = "o"
  board[0][2] = "x"

  for i := 0; i < len(board); i++ {
    fmt.Printf("%s\n", strings.Join(board[i], " "))
  }

}
