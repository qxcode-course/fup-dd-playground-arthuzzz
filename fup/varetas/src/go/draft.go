package main
import "fmt"
func main() {
    var tam1 int
    var tam2 int
    var tam3 int

    fmt.Scan(&tam1, &tam2, &tam3)

    if (tam1 + tam2 <= tam3) || (tam2 + tam3 <= tam1) ||
    (tam3 + tam1 <= tam2) {
        fmt.Println("False")
    } else {
        fmt.Println("True")
    }
}
