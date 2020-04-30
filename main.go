package main

import "test2020/v1/p1"
import "test2020/v1/p2"
import "fmt"

func main() {
    p1.P1 = 100
    fmt.Printf("%d\n", p1.P1)
    fmt.Printf("%d\n", p2.P2)
    fmt.Printf("%s\n", "test1")
}
