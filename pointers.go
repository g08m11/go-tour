// goのmoretypesの学びの内容は全てこのファイルで行う
package main

import "fmt"

func main() {
  // 配列でも後ろに何も無くてもカンマ必要
  names := [4] string {
    "John",
    "Paul",
    "George",
    "Ringo",

  }

  fmt.Println(names)

  a := names[0:2]
  b := names[1:3]
  fmt.Println(a, b)

  b[0] = "XXX"
  fmt.Println(a, b)
  fmt.Println(names)


}
