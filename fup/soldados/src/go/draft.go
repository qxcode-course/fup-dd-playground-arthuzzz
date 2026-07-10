package main
import "fmt"
func main() {
    var nl, nc int
    fmt.Scan(&nl, &nc)

    mt:= make([][]int, nl)

    for i := 0; i < nl; i++ {
        mt[i] = make([]int, nc)
        for j := 0; j < nc; j++ {
            fmt.Scan(&mt[i][j])
        }
    }
    c:= 0
    for i:= 0; i < nc; i++ {
        for j:= 0; j < nl - 1; j ++ {
            if mt[j][i] > mt[j+1][i] {
                c++
            } 
            }
        }
        fmt.Println(c)
    }
