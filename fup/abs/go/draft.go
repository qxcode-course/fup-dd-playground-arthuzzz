
package main
import "fmt"
func main() {
    var A, B int
    fmt.Scan(&A, &B)
    Result := A-B
    if Result < 0 {Result = -Result}

    fmt.Println(Result)
}
