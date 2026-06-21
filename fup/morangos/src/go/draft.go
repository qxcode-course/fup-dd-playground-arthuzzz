package main
import "fmt"
func main() {
    var x1 int
    var y1 int
    var x2 int
    var y2 int
    fmt.Scan(&x1, &y1, &x2, &y2)

    area1:= x1 * y1
    area2:= x2 * y2

    if area1 > area2 {
        fmt.Println(area1)
    } else {
        fmt.Println(area2)
    }
}