package main
import "fmt"

func main() {
    var bento int
    var cebola int
    var qntd int

    fmt.Scan(&bento, &cebola, &qntd)

    animais:= make([]string, qntd)

    for i:= 0; i < qntd; i++ {
        fmt.Scan(&animais[i])
    }

    patas:= 0

    for i := 0; i < qntd; i++ {
        if animais[i] == "v" {
            patas += 4
        }
        if animais[i] == "g" {
            patas += 2
        }
        if animais[i] == "c" {
            patas += 4
        }
    }
    fmt.Println(patas)
    bentov:= (bento - patas) * 1
    cebolav:= (cebola - patas) * 1

    if bentov < 0 {
        bentov = -bentov
    }
    if cebolav < 0 {
        cebolav = -cebolav
    }

    if cebolav < bentov {
        fmt.Println("Cebolinha")
    } else if cebolav > bentov {
        fmt.Println("Chico Bento")
    } else if cebolav == bentov {
        fmt.Println("empate")
    }

}