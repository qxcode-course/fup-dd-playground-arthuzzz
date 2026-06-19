package main
import "fmt"
func main() {
    var p int
    var s int
    var e int
    fmt.Scan(&p, &s, &e)
    fmt.Print("0 ")

    for i:= 300; i <= p; i+=s {
        fmt.Println(i)
        i-=e
        fmt.Print(i, " ")
        s-= 10
        if i >= p {
             fmt.Print("saiu")
            break
        } else if i < 0 {
            fmt.Print("morreu")
            break
        }
    }

}