package main
import "fmt"
func main() {
    var a, b, c int
    var h, l int

    fmt.Scan(&a, &b, &c, &h, &l)

    janela:= h*l

    if (janela >= a*b) || (janela >= b*c) || (janela >= c*a) {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}
