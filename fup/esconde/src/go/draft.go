package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    for i:= 1; i <= N; i+=2 {
            fmt.Println(i)

        }
    for i:= N; i >= 0; i-- {
        if i % 2 == 0 {
        fmt.Println(i)
        }
    }
        
    }

