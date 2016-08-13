package main

import  "fmt"

// deferは実行処理を遅らせる
func main() {
  defer fmt.Println("world")

  fmt.Println("hello")
}
