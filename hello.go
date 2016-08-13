package main

import  "fmt"

// main内でfor実行
// セミコロン省略可(個人的にはコード規約にセミコロン無しを指定したい)
func main() {
  sum := 1
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)
}
