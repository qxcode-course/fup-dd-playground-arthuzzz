package main
import "fmt"
func main() {
    var N int
    var pe string
    fmt.Scan(&N, &pe)

    fmt.Print("[ ")

    for i := 0; i <= 10; i++ {
        if i == N {
            continue
        } 
        if i == 10 {
            fmt.Print("ceu ")
            continue
        }

        fmt.Print(i, pe, " ")
        if pe == "d"{
            pe = "e"
        } else {
            pe = "d"
        }
    }
    fmt.Print("]\n")
}
