package main

import  (
  "fmt"
  "time"
)

// 他言語同様、switch caseは上からの処理
func main() {
  fmt.Print("When's Saturday? ")
  today := time.Now().Weekday()

  switch time.Saturday {
  case today + 0 :
    fmt.Println("Today . ")
  case today + 1 :
    fmt.Println("Tommorow. ")
  case today + 2 :
    fmt.Println("In tow days.")
  default :
    fmt.Println("Too far away.")
  }

}
