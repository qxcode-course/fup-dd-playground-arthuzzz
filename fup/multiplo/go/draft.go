package main
import "fmt"
func main() {
    var a int
    fmt.Scan(&a)

    if (a % 7 == 0) {
        fmt.Println("SIM")
    } else {
        fmt.Println("NAO")
    }

}
