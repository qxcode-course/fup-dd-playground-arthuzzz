package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    fmt.Print("[ ")

    for i:= 0; i <= 10; i++ {
        if i == N {
            continue
        }
        if i==10 {
            fmt.Print("ceu ")
            continue
        }
        fmt.Print(i, " ")
    } 
    fmt.Print("]\n")
}
