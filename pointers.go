// goのmoretypesの学びの内容は全てこのファイルで行う
package main

import "fmt"

// カンマ区切り同一型で定義するの楽
type Vertex struct {
  X, Y int
}

// goは変数名の大文字、小文字の制御厳しい
var  (
  v1 = Vertex{1, 2}
  v2 = Vertex{X: 1}
  v3 = Vertex{}
  p = &Vertex{1, 2}
)


func main() {
  fmt.Println(v1, p, v2, v3)
}
