package main
import "fmt"
func main() {
    var P int
    var D1, D2 int
    fmt.Scan(&D1, D2)
    fmt.Scan(&P)

    soma := (D1 + D2) % 2

    if soma%2 == 0 {
    if P == 0 {
        fmt.Println(0)
    }
    } else if P == 1 {
        fmt.Println("Bob gritou ímpar")
    }


}
