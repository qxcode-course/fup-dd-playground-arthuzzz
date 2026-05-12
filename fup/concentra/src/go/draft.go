package main
import "fmt"
func main() {
    var A, B int
    fmt.Scan(&A,&B)

    fmt.Print("[ ")

	f := true

	for i := A; i <= B; i++ {
		if !f {
			fmt.Print(" ")
		}

		fmt.Print(i)
		fmt.Print(" ")

		fmt.Print(B - (i - A))

		f = false
	}

	fmt.Print(" ]\n")
}