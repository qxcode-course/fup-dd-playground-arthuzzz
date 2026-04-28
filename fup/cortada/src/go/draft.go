package main
import "fmt"
func main() {
    var B int // distância do ponto inicial do corte, na base, para o lado esquerdo da nota.
    var T int // distância do ponto final do corte, no topo, para o lado esquerdo da nota.

    fmt.Scan(&B,&T)

    area := ((B*T)*70/2)
    at:= 70*160
    metade := area/2


    if area == (at - metade) {
        fmt.Println("0")
    } else if area > at {
        fmt.Println("1")
    } else if area < at {
        fmt.Println("2")
    }

 }

