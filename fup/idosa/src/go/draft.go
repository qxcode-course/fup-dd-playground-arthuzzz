package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pessoa struct {
	nome  string
	idade int
	sexo  rune
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	pessoas := make([]Pessoa, n)

	for i := 0; i < n; i++ {
		var sexo string
		fmt.Fscan(in, &pessoas[i].nome, &pessoas[i].idade, &sexo)
		pessoas[i].sexo = rune(sexo[0])
	}

	encontrou := false
	maiorIdade := -1
	nomeMaisIdosa := ""

	for _, p := range pessoas {
		if p.sexo == 'f' {
			if !encontrou || p.idade > maiorIdade {
				encontrou = true
				maiorIdade = p.idade
				nomeMaisIdosa = p.nome
			}
		}
	}

	if encontrou {
		fmt.Println(nomeMaisIdosa)
	} else {
		fmt.Println("nao tem mulher")
	}
}