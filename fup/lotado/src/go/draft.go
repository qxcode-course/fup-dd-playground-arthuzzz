package main
import "fmt"
func main() {
    var C int
//var M int
    fmt.Scan(&C)
    //fmt.Scan(&M)

    soma:= 0

    for {
        var M int
        fmt.Scan(&M)
        soma += M

        if soma == 0 {
                fmt.Println("vazio")
        } else if soma > 0 && soma < C {
                fmt.Println("ainda cabe")
            } else if soma >= C && soma < (C*2) {
                fmt.Println("lotado")
            } else if soma >= (C*2) {
    fmt.Println("hora de partir")
    break
        }
    }
  }
