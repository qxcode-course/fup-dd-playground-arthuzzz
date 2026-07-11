package main
import "fmt"
func main() {
    mt := make([][] int, 3)

    for i := 0; i < 3; i++ {
        mt[i]= make([]int, 3)
    for j:= 0; j < 3; j++ {
        fmt.Scan(&mt[i][j])
    }
}
ref := 0
for i := 0; i < 3; i++ {
    ref += mt[0][i]
}
for i:= 1; i< 3; i++ {
    soma:= 0
    for j:= 0; j < 3; j++ {
        soma += mt[i][j]
    }
    if soma!= ref {
        fmt.Println("nao")
        return
    }
}
for j:= 0; j < 3; j++ {
    soma:= 0
    for i:= 0; i < 3; i++ {
        soma += mt[i][j]
    }
    if soma != ref {
        fmt.Println("nao")
        return
    }
soma = 0

for i:= 0; i < 3; i++ {
    soma += mt[i][2-i]
}
     if soma != ref {
    fmt.Println("nao")
     return
     }
    }
    fmt.Println("sim")
}