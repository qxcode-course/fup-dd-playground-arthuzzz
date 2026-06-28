package main
import "fmt"
func main() {
    var N int

    fmt.Scan(&N)

    cal := make([]int, N)

    for i := 0; i < N; i++ {
        fmt.Scan(&cal[i])
    }

    soma := 0

    for i := 0; i < N; i++ {
        soma += cal[i]
    }
    media:= float64(soma)/float64(N)

    fmt.Printf("%.1f\n", media)


}