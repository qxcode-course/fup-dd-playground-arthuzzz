package main
import "fmt"
func main() {
    var A, B int
    fmt.Scan(&A, &B)

    fmt.Print("[")

    if A > B {
        for ;A > B; A-- {
            fmt.Print(" ", A)
        }
    } else {
        for ;A < B; A++ {
            fmt.Print(" ", A)
        }
    }

    fmt.Print(" ]\n")
}
