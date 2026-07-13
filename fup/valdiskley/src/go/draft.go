package main
import "fmt"
func main() {
    var s string
    var rot int
    fmt.Scan(&s, &rot)

    n:= int(s[0] - 'a')
    resu:= (n+ rot) % 26
    if resu <0 {
        resu+=26
    }
    fmt.Println(string(byte(resu)+ 'a'))
}