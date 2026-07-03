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

    vogais:= ""
    conso:= ""

    for i := 0; i < len(frase); i++ {
        fmt.Scan(frase[i])
        if frase[i] == ' '{
            continue
        }
        if frase[i] == 'a' || frase[i] == 'e' || frase[i] =='i' || frase[i] == 'o' || frase[i] == 'u' {
            vogais += string(frase[i])
        } else {
            conso += string(frase[i])
        }
    }
    fmt.Println(vogais)
    fmt.Println(conso)
}