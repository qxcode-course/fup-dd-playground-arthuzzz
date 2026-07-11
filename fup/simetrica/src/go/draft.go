package main
import "fmt"
func main() {
    mt := make([][]int, 3)

    for i:= 0; i < 3; i++ {
        mt[i] = make([]int, 3)
        for j:= 0; j < 3;  j++ {
            fmt.Scan(&mt[i][j])
        }
    }
    for i := 0; i < 3; i++ {
        for j:= i + 1; j < 3 ; j++ {
            if mt[i][j] != mt[j][i] {
            fmt.Println("nao")
            return
            }
        }
    }
    fmt.Println("sim")
}