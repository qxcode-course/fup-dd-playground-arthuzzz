package main
import "fmt"
func main() {
    var num1, num2 int
    fmt.Scan(&num1, &num2)

    divint := num1/num2
    resto := num1%num2
    DivFloat := float64(num1) / float64(num2)

    fmt.Println(divint)
    fmt.Println(resto)
    fmt.Printf("%.2f\n", DivFloat)


}
