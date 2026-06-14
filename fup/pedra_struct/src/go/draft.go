package main

import (
	"bufio"
	"fmt"
	"os"
)

type Jogada struct {
	p1, p2 int
}

// retorna se a jogada é válida e sua pontuação
func calc_pontuacao(jogada Jogada) (bool, int) {
	if jogada.p1 < 10 || jogada.p2 < 10 {
		return false, 0
	}

	diff := jogada.p1 - jogada.p2
	if diff < 0 {
		diff = -diff
	}

	return true, diff
}

func procurar_melhor_jogada(jogadas []Jogada) int {
	melhorIndice := -1
	melhorPontuacao := 0

	for i, jogada := range jogadas {
		valida, pontuacao := calc_pontuacao(jogada)

		if !valida {
			continue
		}

		if melhorIndice == -1 || pontuacao < melhorPontuacao {
			melhorIndice = i
			melhorPontuacao = pontuacao
		}
	}

	return melhorIndice
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	jogadas := make([]Jogada, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &jogadas[i].p1, &jogadas[i].p2)
	}

	vencedor := procurar_melhor_jogada(jogadas)

	if vencedor == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(vencedor)
	}
}