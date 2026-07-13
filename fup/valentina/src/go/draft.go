package main
import "fmt"
func main() {
    var s1, s2 string
    var op string
    fmt.Scan(&s1)
    fmt.Scan(&op)
    fmt.Scan(&s2)

    var resu int

    c1:= s1[0]
    c2:= s2[0]

    n1:= int(c1 -'a')
    n2:= int(c2 -'a')

    if op == "+" {
    resu = (n1+n2) % 26
    } else {
        resu = (n1 - n2 + 26) % 26
    }
    fmt.Println(string(byte(resu)+'a'))
}