package p2

import "test2020/v1/p1"
import "fmt"

const P2 = 45

func init() {
  p1.P1 = 72
  fmt.Println("p2 init")
}
