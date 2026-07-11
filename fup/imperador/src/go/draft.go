package main
import "fmt"
func main() {
    var n int 
    fmt.Scan(&n)
    mt := make([][]string, n)

    ll := -1
    lc:= -1

    for i := 0; i < n; i++ {
        mt[i] = make([]string, n)
        for j:= 0; j < n; j++ {
            fmt.Scan(&mt[i][j])
            if mt[i][j] == "L" {
                ll = i 
                lc = j
            }
        }
    }
    glad:= 0
    cond := 0

for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
        if ll != -1 && (i == ll || j == lc) {
            continue
        }

        if mt[i][j] == "G" {
            glad += 2
        } else if i+j == n-1 {
            cond += 2
        } else {
            cond++
        }
     }
}
if glad > cond {
    fmt.Println("Gladiadores")
} else if cond > glad {
        fmt.Println("Condenados a morte")
    } else {
        fmt.Println("Ninguem")
    }
}
