package main
import "fmt"
func main() {
    var total int
    fmt.Scan(&total)

    var qntd int
    fmt.Scan(&qntd)

    fig:= make([]int, qntd)

    for i := 0; i < qntd; i++ {
        fmt.Scan(&fig[i])
    }
    fmt.Print("[ ")

    for i := 1; i < qntd; i++ {
        if fig[i] == fig[i-1] {
            fmt.Print(fig[i], " ")
        }
    }
    fmt.Print("]")
    fmt.Println()

    perc:= 0

    fmt.Print("[ ")

    for u:= 1; u <= total; u++  {
        for perc < qntd && fig[perc] < u {
               perc++
        }
        if perc >= qntd || fig[perc] != u {
            fmt.Print(u, " ")
        }
    }
    fmt.Println("]")
}