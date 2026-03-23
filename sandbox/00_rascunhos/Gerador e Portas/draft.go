package main
import "fmt"
func main() {
    Fonte := true
    if Fonte == false{
        fmt.Println("O gerador está desativado!")
    } else if Fonte == true{
        fmt.Println("O gerador está ativado!")
    }


    Status := true
    if Fonte == false{
        fmt.Println("A lâmpada está apagada,")
    } else if Status == false{
    fmt.Println("A lâmpada está apagada,")
}    else if Status == true{
    fmt.Println("A lâmpada está acesa,")}

Statusd := false

if Statusd == false{
    fmt.Println("a porta está fechada")
} else if Statusd == true{
    fmt.Println("a porta está aberta")
}

PCStatus := true
if Fonte == false{
fmt.Println("e o PC está desligado")
} else if PCStatus == false{
    fmt.Println("e o PC está desligado")
} else if PCStatus == true{
    fmt.Println("e o PC está ligado")
}
}