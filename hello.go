package main

import  "fmt"

// deferは実行処理を遅らせる
// この場合、fmtの処理はdeferによりLIFOのロジックとなってしまう。
func main() {
  defer fmt.Println("counting")
  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }
  fmt.Println("done")
}
