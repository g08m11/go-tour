// goのmoretypesの学びの内容は全てこのファイルで行う
package main

import "fmt"

var pow = []int{1, 2, 4, 6, 8, 16, 32, 64, 128}

func main() {
  for i, v := range pow {
    fmt.Printf("2**%d = %d\n", i, v)
  }
}
