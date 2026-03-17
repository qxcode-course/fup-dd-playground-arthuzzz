package main
import "fmt"
func main() {
    var num1, num2 int

    fmt.Scan(&num1)
    fmt.Scan(&num2)

    media := float64(num1 + num2)/2

    fmt.Printf("%.1f\n", media)
}
