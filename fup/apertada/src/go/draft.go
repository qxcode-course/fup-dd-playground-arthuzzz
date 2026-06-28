package main
import "fmt"
func main() {
    var x, y, z, a, b int
    fmt.Scan(&x, &y ,&z, &a, &b)

    trinta:= make([]int, 30)

    for i:= 1; i <= len(trinta); i++ {
        if i == x || i == y || i == z || i == a || i == b {
            fmt.Println(i)
            break
        }
    }
}
