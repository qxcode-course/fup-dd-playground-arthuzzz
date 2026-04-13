package main
import "fmt"
func main() {
    var A, B int
    fmt.Scan(&A, &B)

    fmt.Print("[")


    for true {
        if A > B {
            break
        } 
        if A % 2 == 0 {
            A++
            continue
        } 
        fmt.Print(" ", A)
        A++
    }
    fmt.Print(" ]\n")
}
