package main
import "fmt"
func main() {
    var P int
    var D1, D2 int
    fmt.Scan(&P, &D1, &D2)

    soma := (D1 + D2) % 2


    if P == 0 {
        if soma == 0 {
            fmt.Println(0)
        } else {
            fmt.Println(1)
    }
    } else {
        if soma == 0 {
        fmt.Println("1")
    } else {
        fmt.Println(0)
    }

}
}