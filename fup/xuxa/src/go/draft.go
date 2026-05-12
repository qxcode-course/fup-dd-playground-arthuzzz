package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	leitor := bufio.NewReader(os.Stdin)

	frase, _ := leitor.ReadString('\n')

	if len(frase) > 0 && frase[len(frase)-1] == '\n' {
		frase = frase[:len(frase)-1]
	}

	caracteres := []rune(frase)

	for i := len(caracteres) - 1; i >= 0; i-- {
		fmt.Print(string(caracteres[i]))
	}

	fmt.Println()
}