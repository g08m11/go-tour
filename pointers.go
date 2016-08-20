// goのmoretypesの学びの内容は全てこのファイルで行う
package main

import "fmt"



func main() {
  // array宣言もRubyと似ててイメージつきやすい
  // 配列のみ指定すると配列型で値を返す
  // := も使えるため型と値設定を同時に可能
  var a[2]string
  a[0] = "Hello"
  a[1] = "World"
  fmt.Println(a[0], a[1])
  fmt.Println(a)
  primes := [6]int{2, 3, 5, 7, 11, 13}
  fmt.Println(primes)
}
