package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	fator := 2
	contagem := 0

	for N != 1 {

		if N%fator == 0 {
			N = N / fator
			contagem++

		} else {

			if contagem > 0 {
				fmt.Println(fator, contagem)
			}

			fator++
			contagem = 0
		}
	}
	if contagem > 0 {
		fmt.Println(fator, contagem)
	}
}