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

    frase:= ler()
	trecho := ler()

	contador:= 0

    for i := 0; i <= (len(frase) - len(trecho)); i++ {
        if frase[i:i+len(trecho)] == trecho {
			contador++
		}
    }
	fmt.Println(contador)
}