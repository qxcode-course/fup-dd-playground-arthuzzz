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

    for i := 0; i < q; i++ {
        fmt.Scan(&pos[i].x, &pos[i].y)
    }

    cab:= make([] Ponto, q)

    switch d {
    case "L":
        cab[0].x = pos[0].x - 1
        cab[0].y = pos[0].y

    case "R" :
        cab[0].x = pos[0].x + 1
        cab[0].y = pos[0].y

    case "U":
        cab[0].x = pos[0].x
        cab[0].y = pos[0].y - 1

    case "D" :
        cab[0].x = pos[0].x
        cab[0].y = pos[0].y + 1
    }
    for i:= 1; i < q ; i++ {
        cab[i] = pos[i-1]
    }

    for i:= 0; i < q; i++ {
        fmt.Println(cab[i].x, cab[i].y)
    }


    
}