// goのmoretypesの学びの内容は全てこのファイルで行う
package main

import "fmt"

func main() {
  // array宣言もRubyと似ててイメージつきやすい
  // 範囲指定の場合はRubyの「..」ではなく「:」なのがpoint
  primes := [6] int {2, 3, 5, 7, 11, 13}

  var s [] int = primes[1:4]
  fmt.Println(s)
}
