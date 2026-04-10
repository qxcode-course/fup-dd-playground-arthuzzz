package main

import (
	"fmt"
	"math"
)
func main() {
    var delta float64
    var a float64
    var b float64
    var c float64
    
     var Raiz1, Raiz2 float64

    fmt.Scan(&delta, &a, &b, &c)

    delta = (math.Pow(b, 2)) - (4*a*c)
    Raiz1 = ((-b) + (math.Sqrt(delta))) / 2
    Raiz2 = ((-b) - (math.Sqrt(delta))) / 2

    if delta >= 0 {
        fmt.Printf("%.2f\n",Raiz1)
        fmt.Printf("%.2f\n",Raiz2)
    } else if delta == 0 {
        fmt.Printf("%.2f\n",Raiz1)
    } else {
        fmt.Println("nao ha raiz real")
    }



}
