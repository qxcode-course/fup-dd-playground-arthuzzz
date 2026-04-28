package main
import "fmt"
func main() {
    var valor, parcelas float64
    fmt.Scan(&valor, &parcelas)

    juros := (((parcelas * 5) - 5)/100)
    total := (valor + (valor * juros))

    if parcelas == 1 {
        fmt.Printf("%.2f\n", valor)
        fmt.Printf("%.2f\n", valor)
    } else {
        fmt.Printf("%.2f\n", (total)/parcelas)
        fmt.Printf("%.2f\n", total)
    }



}
