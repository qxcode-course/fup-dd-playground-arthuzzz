package main
import "fmt"
func main() {
    var P int
    var D1, D2 int
    fmt.Scan(&D1, D2)
    fmt.Scan(&P)

    soma := (D1 + D2)

    if soma%2 == 0 {
    if P == 0 {
        fmt.Println(0)
    } else {
        fmt.Println(1)
    } else {
        if P == 1 {
            fmt.Println(1)
        } else {
            fmt.Println(0)
        }
    }


}
