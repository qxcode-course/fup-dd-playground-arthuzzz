package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	leds := [10]int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}

	for i := 0; i < n; i++ {
		var numero string
		fmt.Fscan(in, &numero)

		total := 0
		for _, c := range numero {
			total += leds[c-'0']
		}

		fmt.Printf("%d leds\n", total)
	}
}