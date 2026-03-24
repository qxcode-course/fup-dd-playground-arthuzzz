package main
import "fmt"
func main() {
    var kmh, tempo, consumo int
    fmt.Scan(&kmh, &tempo, &consumo)


    T_H := tempo/60
    DP := kmh*T_H
    DF := DP/consumo
}
