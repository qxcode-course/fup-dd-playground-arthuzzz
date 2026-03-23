package main
import "fmt"
func main() {
    var T_c float64
    fmt.Scan(&T_c)

    T_f := 1.8*T_c+32

    fmt.Printf("%.6f\n", T_f)

}
