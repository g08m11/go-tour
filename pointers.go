// goのmoretypesの学びの内容は全てこのファイルで行う
package main

import "fmt"

// ruby同様にarrayへの追加をappendで行うので理解しやすい
func main() {
  var s []int
  printSlice(s)

  s = append(s, 0)
  printSlice(s)

  s = append(s, 1)
  printSlice(s)

  s = append(s, 2, 3, 4)
  printSlice(s)
}

func printSlice(s []int ) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
