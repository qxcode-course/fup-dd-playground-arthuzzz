package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    fmt.Print("[ ")

    for i:= 0; i <= 10; i++ {
        if i == 10 {
            fmt.Print( "ceu")
        }
        
        if i == N {
            continue
        }
        
        fmt.Print(i, " ")
    }
    fmt.Print("]")
}
