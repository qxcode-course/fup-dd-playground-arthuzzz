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
    e:= ler()
    key:= ler()

    for i := 0; i < len(e); i++ {
        k:= key[i%len(key)] - '0' 

        n := e[i] ^ k

        fmt.Print(string(n))
    }
    fmt.Println()
}