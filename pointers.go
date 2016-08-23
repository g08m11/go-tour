// goのmoretypesの学びの内容は全てこのファイルで行う
package main

import "fmt"

func main() {
  pow := make([]int, 10)
  for i := range pow {
    pow[i] = 1 << uint(i) // == 2**i
  }

  // swiftみたいにアンダーバーで代入も可能
  // あまりこの書き方分かりづらいから未経験の人がいる場合は書きたくない(長過ぎる名前なら別)
  // コーディング規約に明記したい
  for _, value := range pow {
    fmt.Printf("%d\n", value)
  }
}
