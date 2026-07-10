package main
import "fmt"
func main() {
    var j1, j2, j3, j4 int
    fmt.Scan(&j1, &j2, &j3, &j4)

    dedos:= j1 + j2 + j3 + j4

    if dedos == 0 {
        fmt.Println("nenhum")
    } else if dedos == 4 || dedos / 2 == 4 {
        fmt.Println("jog4") 
    }else if dedos == 2 || dedos / 7 == 2 {
        fmt.Println("jog2")
    } else if dedos == 1 || dedos - 4 == 1 {
        fmt.Println("jog1")
    } else if dedos == 3 {
        fmt.Println("jog3")
    }
}