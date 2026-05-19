package main
import "fmt"
func main() {
    var C, M int
    fmt.Println(&C, &M)

    for i := M; i <= (C*2); i++ {
    if i == 0 {
        fmt.Println("vazio")
        continue
    }
    if i < C {
        fmt.Println("ainda cabe")
        continue
  }
  if M == (C*2) {
    fmt.Println("hora de partir")
    break
}
}
}