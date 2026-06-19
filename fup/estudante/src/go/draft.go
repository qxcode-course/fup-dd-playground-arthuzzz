package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Aluno struct {
	Nome  string
	N1    float64
	N2    float64
	N3    float64
	Media float64
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	alunos := make([]Aluno, n)

	for i := 0; i < n; i++ {
		nome, _ := reader.ReadString('\n')
		nome = nome[:len(nome)-1] // remove '\n'

		var n1, n2, n3 float64
		fmt.Fscanln(reader, &n1, &n2, &n3)

		alunos[i] = Aluno{
			Nome:  nome,
			N1:    n1,
			N2:    n2,
			N3:    n3,
			Media: (n1 + n2 + n3) / 3.0,
		}
	}

	sort.Slice(alunos, func(i, j int) bool {
		return alunos[i].Media > alunos[j].Media
	})

	for i, a := range alunos {
		fmt.Printf("%d %s %.2f %.2f %.2f %.2f\n",
			i,
			a.Nome,
			a.Media,
			a.N1,
			a.N2,
			a.N3)
	}
}