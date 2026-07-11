package main
import "fmt"
func main() {
    var a, g, ra, rg float64
    fmt.Scan(&a, &g, &ra, &rg)

    ca := a / ra
    cg := g / rg 

    if ca < cg {
        fmt.Println("A")
    } else {
        fmt.Println("G")
    }

}