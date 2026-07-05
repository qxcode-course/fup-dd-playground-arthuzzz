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
func vogal(c byte) bool {
        return c == 'a' || c =='e' || c == 'i' ||
        c == 'o' || c == 'u'
    }

func main() {
    frase:= ler()
    resu:= ""

    for i := 0; i < len(frase); i++ {
        if frase[i] == ' ' && i > 0 && i < len(frase)-1 {
            if vogal(frase[i-1]) && vogal(frase[i+1]) {
                continue
            }
        }

        if vogal(frase[i]) {
            for len(resu) >= 3 &&
            vogal(resu[len(resu) -1]) && vogal(resu[len(resu)-2]) && vogal(resu[len(resu)-3]) {
                resu = resu[:len(resu)-2] + resu[len(resu)-1:]
            }
        }
        resu += string(frase[i])
        if len(resu) >= 3 &&
            vogal(resu[len(resu) -1]) && vogal(resu[len(resu)-2]) && vogal(resu[len(resu)-3]) {
                resu = resu[:len(resu)-2] + resu[len(resu)-1:]
            }
    }
     fmt.Println(resu)
    }
