package main
import "fmt"
func main() {
    var p int
    var s int
    var e int
    fmt.Scan(&p, &s, &e)
    fmt.Print("0 ")

    for i:= s; i <= (p+s); i+=s {
         if i >= p {
             fmt.Print("saiu\n")
            break
        }
        fmt.Println(i)
        i-=e
        fmt.Print(i, " ")
        s-= 10
        if i < 0 {
            fmt.Print("morreu\n")
            break
        }
    
    }

}