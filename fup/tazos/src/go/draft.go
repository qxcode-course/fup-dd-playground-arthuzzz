package main
import "fmt"
func main() {
    var qntd int
    fmt.Scan(&qntd)

    fig:= make([]int, qntd)

    repete:= 0
    for i:= 0; i < qntd; i++ {
        if fig[i] == fig[i-1] {
            repete++
            fmt.Print(fig[i])
        }
    }
}