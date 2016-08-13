package main

import (
  "fmt"
  "math"
)

// 型まで意識しての出力はprintfとなるため注意
func main() {
  // 変数宣言 var枠でひとまとめ
  // 短縮の「:=」で書くのもあり。値が決まっていたら。
  // こっちの方がなんかRubyっぽくて好き
  x := 3
  y := 4
  f := math.Sqrt(float64(x*x + y*y))
  z := uint(f)

  fmt.Println(x, y, z)

}
