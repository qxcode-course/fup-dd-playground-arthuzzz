package main
import "fmt"
func main() {
    var A int
    var B int

    fmt.Scan(&A, &B)

    for ;A < B; A++ {
        fmt.Println(A)
    }
}
