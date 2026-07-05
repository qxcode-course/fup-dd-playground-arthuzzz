package main
import "fmt"

type restaurant struct {
    nome string
    nota int
    }

func ganha(resta []restaurant) string  {

    melhor:= resta[0]
    
    for i:= 0; i < len(resta); i++ {
        if resta[i].nota > melhor.nota {
            melhor = resta[i]
        } else if resta[i].nota == melhor.nota && resta[i].nome < melhor.nome  {
            melhor = resta[i]
        }
    }
    return melhor.nome
}

func main() {
    var n int
    fmt.Scan(&n)

    resta:= make([]restaurant, n)

    for i:= 0; i < n; i++ {
        fmt.Scan(&resta[i].nome, &resta[i].nota)
    }

    fmt.Println(ganha(resta))

}