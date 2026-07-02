package main
import "fmt"
func main() {
    var P, N int
    fmt.Scan(&P, &N)

    contador:= 0
    ps:= make([] int, N)

    for i := 0; i < N; i++ {
        fmt.Scan(&ps[i])
    }

    for i:= 0; i < N; i++ {
        if ps[i] == P {
            contador++
        }
    }
    fmt.Println(contador)
}