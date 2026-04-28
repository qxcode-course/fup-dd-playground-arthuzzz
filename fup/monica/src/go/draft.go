package main
import "fmt"
func main() {
    var M int
    var A int
    var B int

    fmt.Scan(&M, &A, &B)

    idade := M - (A+B)

    if idade > A && idade > B {
        fmt.Println(idade)
    } else if A > B && A > idade {
            fmt.Println(A)
    } else {
            fmt.Println(B)
    }
 }


