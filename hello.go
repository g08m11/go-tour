package main

import  "fmt"

// main内でfor実行
func main() {
  sum := 1
  for ; sum < 1000; {
    sum += sum
  }
  fmt.Println(sum)
}
