package main

import (
	"fmt"
	//"math/rand"
)
func main() {
    const pedra int = 0
    const papel int = 1
    const tesoura int = 2

    var jog1, jog2 int

    p1:= 0
    p2:= 0
    for r:= 1; r <= 5; r++ {
        fmt.Println("# JOKENPÔ #")
        fmt.Printf("jogador 1: %d | Jogador 2: %d\n", p1, p2)
        fmt.Printf("Round: %d/5\n\n", r)
    
    
    fmt.Print("Jog1: digite 0 (pedra), 1 (papel), tesoura(2) : ")
    fmt.Scan(&jog1)

    fmt.Print("Jog2: digite 0 (pedra), 1 (papel), tesoura(2): ")
    fmt.Scan(&jog2)

    if jog1 == jog2 {
        fmt.Println("empate")
    } else if (jog1 == pedra && jog2 == tesoura) ||
    (jog1 == papel && jog2 == pedra) ||
    (jog1 == tesoura && jog2 == papel) {
        fmt.Println("Jogador 1 venceu!")
    } else {
        fmt.Println("Jogador 2 venceu!")
    }
}
    }



