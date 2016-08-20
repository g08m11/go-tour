// goのmoretypesの学びの内容は全てこのファイルで行う
package main

import "fmt"

func main() {

  var s [] int
  fmt.Println(s, len(s), cap(s))
  if s == nil {
    fmt.Println("Hey! you fuck off ! no value ! nil!")
  }
}
