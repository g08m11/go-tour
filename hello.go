package main

import  (
  "fmt"
  "math"
)

// セミコロンイケる場合といけない場合の区別しづらい。
// ある条件であれば式を代入するという時はチームで話あう必要がある。
func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  return lim
}

// 後述がないのにカンマを入れないといけないのも分かりづらい
func main() {
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )
}
