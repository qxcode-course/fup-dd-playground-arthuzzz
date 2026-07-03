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
	texto := ler()

	var inicio, quantidade int
	fmt.Scan(&inicio)
	fmt.Scan(&quantidade)

	if inicio < 0 || inicio >= len(texto) || quantidade < 0 {
		fmt.Print("")
		return
	}

	fim := inicio + quantidade

	if fim > len(texto) {
		fim = len(texto)
	}

	fmt.Println(texto[inicio:fim])
}