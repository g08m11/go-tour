// goのmoretypesの学びの内容は全てこのファイルで行う
package main

import "fmt"

func main() {

  s := [] int { 2, 3, 5, 7, 11, 13 }
  fmt.Println(s)

  // 1から4までを取得するそのため、取得個数は3
  s = s[1:4]
  fmt.Println(s)

  // 0から2個の要素を前から取得する。余ったものは切り捨て
  s = s[:2]
  fmt.Println(s)

  // 0から1個の要素を前から取得する。余ったものは切り捨て
  s = s[1:]
  fmt.Println(s)
}
