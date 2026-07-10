package main
import "fmt"
func main() {
    var N int 
    fmt.Scan(&N)

    mt:= make([][]int, N)

    for i:= 0; i < N; i++ {
        mt[i] = make([]int, N)
        for j:= 0; j < N; j++ {
            fmt.Scan(&mt[i][j])
        }
    }

    maior:= 1
    in:= 0

    for i:= 0; i < N; i++ {
        soma:= 0

        for j:= 0; j < N; j++ {
            soma+= mt[j][i] * mt[j][i]
        }
        if soma > maior {
            maior = soma
            in = i
        }
    }
    fmt.Println(in)


}