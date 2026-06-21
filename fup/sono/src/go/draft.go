package main
import "fmt"
func main() {
    var h, m, s int
    var hf, mf, sf int
    fmt.Scan(&h, &m, &s, &hf, &mf, &sf)

    resth:= hf - h
    restm:= mf - m
    rests:= sf - s

    if resth < 0 {
        restm - resth
    }
}