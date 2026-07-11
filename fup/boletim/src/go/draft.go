package main
import "fmt"
func main() {
    mt:= make([][]int, 2)

    soma:= 0 

    for i := 0; i < 2; i++ {
        mt[i] = make([]int, 3)
        for j:= 0; j < 3; j++ {
            fmt.Scan(&mt[i][j])
            soma += mt[i][j]
        }
    }
    fmt.Println(soma)
}