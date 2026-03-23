package main

import (
	"fmt"
)

func main() {
    var quan1, quan2, quan3 int
    var prod1, prod2, prod3 float64
    var money float64

    fmt.Scan(&quan1, &quan2, &quan3)
    fmt.Scan(&prod1, &prod2, &prod3)
    fmt.Scan(&money)

    Gasto := float64(quan1)*prod1 + float64(quan2)*prod2 + float64(quan3)*prod3

    Troco := money - Gasto

    fmt.Printf("%.2f\n", Troco)
}
