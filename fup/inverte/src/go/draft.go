package main
import "fmt"
func main() {
    var a string
    fmt.Scan(&a)

    x:= a[0]

    if x >= 'A' && x <= 'Z' {
        x = x + 32
    } else if x  >= 'a' && x <= 'z' {
        x = x- 32
    }
    fmt.Printf("%c\n", x)
}