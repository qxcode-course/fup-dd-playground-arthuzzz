package main
import "fmt"
func main() {
    var num int
    fmt.Scan(&num)
    quad := 0

    for i := quad; i >= num; i++ {
        if i * i == num {
            fmt.Println("sim")
            break
        } if i * i == !num {
            fmt.Println("nao")
        }
    }

}