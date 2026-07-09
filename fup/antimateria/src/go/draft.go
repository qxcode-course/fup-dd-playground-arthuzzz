package main
import "fmt"
func main() {
    var a string
    var b string
    fmt.Scan(&a, &b)

    lim := len(a)
    num := 0

    if len(b) < lim {
        lim = len(b)
    }

    for i := lim; i >= 1; i-- {
        fim:= a[len(a) - i:]
        ini := b[:i]

        if fim == ini {
            num = i
            break
        }
    }

    fmt.Println(a[:len(a) - num] + b[num:])
}