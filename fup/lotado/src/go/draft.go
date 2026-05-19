package main
import "fmt"
func main() {
    var C, M int
    fmt.Scan(&C, &M)

    for i := M; i <= (C*2); i++ {
        if i == (C*2) {
    fmt.Println("hora de partir")
    break
        } else if i == 0 {
        fmt.Println("vazio")
    } else if i < C {
        fmt.Println("ainda cabe")
  } else if i >= C {
    fmt.Println("lotado")
  }
}
}
