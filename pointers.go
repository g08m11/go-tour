// goのmoretypesの学びの内容は全てこのファイルで行う
package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

// ドットで呼び出せる感じはRubyと似ていて分かりやすい
// structのフィールドは、structのポインタを通してアクセスすることも出来る
func main() {
  v := Vertex{1, 2}
  p := &v
  p.X = 1e9
  fmt.Println(v)
}
