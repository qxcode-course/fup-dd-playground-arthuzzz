package main
import "fmt"
func main() {
    var a, b int
    var soma int = 0
    fmt.Scan(&a, &b)

    if a > b {
        fmt.Println("invalido")
    } else {
        for ;a > b; a++ {
            if a % 2 == 0{
                soma += a
            
            fmt.Println(soma)
        }
    }
}
}
