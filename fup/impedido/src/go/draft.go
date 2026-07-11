package main
import "fmt"
func main() {
    var l, r, d int
    fmt.Scan(&l, &r, &d)

    if l > 50 && l < r && r > d {
        fmt.Println("S")
    } else if l == d {
        fmt.Println("N")
    } else {
        fmt.Println("N")
    }
}