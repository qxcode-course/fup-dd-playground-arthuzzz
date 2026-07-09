package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ler() string {
    leitor:= bufio.NewReader(os.Stdin)
    texto, _:= leitor.ReadString('\n')
    return strings.TrimSpace(texto)
}

func vogal(c byte) bool {
    return c == 'a' || c == 'e' ||
    c == 'i' || c == 'o' || c == 'u' ||
    c == 'A' || c == 'E' || c == 'I' ||
    c == 'O' || c == 'U'

}

func main() {
    frase:= ler()
    words:= strings.Fields(frase)

    for p := 0; p < len(words); p++ {
        word := words[p]
        pos := -1

        for i := 0; i < len(word)-1; i ++ {
            if vogal(word[i]) && !vogal(word[i+1]) {
                pos = i
                break

        }
    }

    otra:= pos != -1 && pos < len(word)-1

    if pos != -1 && otra {
        sil := word[:pos+1]
        fmt.Print(sil)
        fmt.Print(sil)
        fmt.Print(word) 
    } else {
        fmt.Print(word)
    }
    if p < len(words)-1 {
        fmt.Print(" ")
    }
}
fmt.Println()
}