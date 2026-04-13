package main
import "fmt"
func main() {
    var M int
    var a int
    var b int
    fmt.Scan(&M, &a, &b)

    c := M - (a + b)

    if a > c && a > b {
        fmt.Println(a)
    } else if b > a && b > c {
        fmt.Println(b)
    } else {
        fmt.Println(c)
    }
    } 
