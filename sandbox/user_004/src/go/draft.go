package main
import "fmt"
func main() {
    var a, c int
    var b string
    fmt.Scan(&a, &b, &c)

    if b == "+" {
        fmt.Println("=", a+c)
    } else if b == "-" {
        fmt.Println("=", a-c)
    } else if b == "*" {
        fmt.Println("=", a*c)
    } else if b == "/" {
        fmt.Println("=", a/c)
    }
}