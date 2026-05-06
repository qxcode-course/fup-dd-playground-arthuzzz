package main
import "fmt"
func main() {
    var ini, fim int
    fmt.Scan(&ini, &fim)

    for i:= ini; i <= fim; i++ {
        if (i % 3 == 0) && (i % 5 == 0) {
            fmt.Println("zigzag")
            continue
        } else if (i % 5 == 0) {
            fmt.Println("zag")
            continue
        } else if  (i % 3 == 0) {
            fmt.Println("zig")
            continue
        }
        fmt.Println(i)
        
    } 
    
}
