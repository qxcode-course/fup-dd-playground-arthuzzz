package main
import "fmt"
func main() {
    type Ponto struct {
        x int
        y int
    }

    var q int
    var d string
    fmt.Scan(&q, &d)

    pos:= make([]Ponto, q)

    for i := 0; i < 1; i++ {
        fmt.Scan(&pos[i].x, &pos[i].y)
    }


    
}