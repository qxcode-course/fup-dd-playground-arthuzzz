package main

import (
	"fmt"
	"math"
)
func main() {
    var op string
    var num float64

    fmt.Scan(&op, &num)

    if op == "c" {
        fmt.Println(math.Ceil(num))
    } else if op == "f" {
        fmt.Println(math.Floor(num))
    } else if op == "r" {
        fmt.Println(math.Round(num))
    }

}
