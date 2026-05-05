package main
import "fmt"
func main() {
    var novo, qnt int
    fmt.Scan(&novo, &qnt)

    for i:= 0; i < qnt; i++ {
        idade:= novo + (i * 2)
        fmt.Println(idade)
    } 
}
