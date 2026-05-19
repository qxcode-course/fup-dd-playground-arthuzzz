package main
import "fmt"
func main() {
    var p, s, e int
    fmt.Scan(&p, &s, &e)

    i:= 0
    fmt.Print(i, " ")

    for {
        i += s

         if i >= p {
            fmt.Println("saiu")
            break
         }

            fmt.Print(i, "\n")
    i -= e
    fmt.Print(i, " ")
}
}
