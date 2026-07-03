package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ler() string {
	leitor := bufio.NewReader(os.Stdin)
	texto, _ := leitor.ReadString('\n')
	return strings.TrimSpace(texto)
}
func main() {
    texto:= ler()

    resultado:= ""

    for i := 0; i < len(texto); i++ {

    if texto[i] >= 'A' && texto[i] <= 'Z' {
        resultado += string(texto[i] + 32)
    } else if texto[i] >= 'a' && texto[i] <= 'z' {
        resultado += string(texto[i] - 32)
    } else {
        resultado += string(texto[i])
    }
}
fmt.Println(resultado)
}