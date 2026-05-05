package main
import "fmt"
func main() {
    var c int
    var a int
    fmt.Scan(&c, &a)

    capacidade:= c-1
    viagens:= a/capacidade
    
    if c < a {
        if (viagens/2 == 0) {
        fmt.Println(viagens+1)
        } else {
            fmt.Println(viagens)
        }
    } else if c > a {
        fmt.Println(capacidade/a)
    }
}

