package main
import "fmt"
func main() {

    var num1, num2 int
    var quociente, resto int

    fmt.Scan(&num1)

    fmt.Scan(&num2)

    quociente = num1/num2
    resto = num1 % num2

    fmt.Println(quociente, resto)
}
