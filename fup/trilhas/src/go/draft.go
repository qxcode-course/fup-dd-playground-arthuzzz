package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    best:= 1
    menos:= 1000000

    for t:= 1; t <= n; t++ {
        var m int
        fmt.Scan(&m)

    alt:= make([]int, m)

    for i:= 0; i < m; i++ {
        fmt.Scan(&alt[i])
    }
    
    e1 := 0
    e2:= 0

    for i := 0; i < m-1; i++ {
        if alt[i+1] > alt[i] {
            e1 += alt[i+1] - alt[i]
        }
    }


    for  i := m - 1; i  > 0; i-- {
        if alt[i-1] > alt[i] {
            e2 += alt[i-1] - alt[i]
        }
        }
        es:= e1

        if e2 < es {
            es = e2
        }
        if es < menos {
            menos = es
            best = t
        }
        }
        fmt.Println(best)
    }
