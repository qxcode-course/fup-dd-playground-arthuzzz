package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
var leitor = bufio.NewReader(os.Stdin)
func ler() string {
	texto, _ := leitor.ReadString('\n')
	return strings.TrimSpace(texto)
}

func main() {
    texto:= ler()
    ant:= ler()
    nov:= ler()

    resu:= ""

    for i:= 0; i < len(texto); {
        if i+len(ant) <= len(texto) &&
        texto[i:i+len(ant)] == ant {
            resu += nov
            i+= len(ant)
        } else {
            resu += string(texto[i])
            i++
        }
    }
    fmt.Println(resu)
}