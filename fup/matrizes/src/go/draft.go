package main
import "fmt"
func main() {
    var nl, nc int
    fmt.Scan(&nl, &nc)

    a:= make([][] int, nl)
    b:= make([][] int, nl)

    for i := 0; i < nl; i++ {
        a[i] = make([]int, nc)
        for j := 0; j < nc; j++ {
            fmt.Scan(&a[i][j])
        }
    }

    for i := 0; i < nl; i++ {
        b[i] = make([]int, nc)
        for j:= 0; j < nc; j++ {
            fmt.Scan(&b[i][j])
        }
    }


    for i:= 0; i < nl; i++ {
        fmt.Print("[ ")
        for j:= 0; j < nc; j++ {

        fmt.Print((a[i][j] + b[i][j]))
        if j < nc- 1 {
            fmt.Print(" ")
        }
        }
        fmt.Println(" ]")
    }
}