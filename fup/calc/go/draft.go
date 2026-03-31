package main
import "fmt"
func main() {
    var a, b int
    var c string
    
    fmt.Scan(&a, &b, &c)
    
    switch c {
    case "+":
        fmt.Println(a+b)
    case "-":
        fmt.Println(a-b)
    case "*":
        fmt.Println(a*b)
    case "/":
        fmt.Println(a/b)
    }


    

}
