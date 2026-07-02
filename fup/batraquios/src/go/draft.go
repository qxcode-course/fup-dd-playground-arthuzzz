package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    vetn:= make([]int, n)

    for i:= 0; i < n; i++ {
    fmt.Scan(&vetn[i])
    }

    var m int
    fmt.Scan(&m)

    vetm:= make([]int, m)

    for i := 0; i < m; i ++ {
        fmt.Scan(&vetm[i])
    }

    for i:= 0; i < n; i++ {
        encontrou:= false

        for u := 0; u < m; u++ {
            if vetn[i] == vetm[u] {
                encontrou = true
                break
            }
        }
        if !encontrou {
            fmt.Println("nao")
            return
        }
     }
     fmt.Println("sim")
    }
