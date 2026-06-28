package main
import "fmt"
func main() {
    var h, m, s int
    var hf, mf, sf int
    fmt.Scan(&h, &m, &s, &hf, &mf, &sf)

    rh:= hf - h
    rm:= mf - m
    rs:= sf - s

    if rh < 0 {
        
    }
    }

    fmt.Printf("%02d", rh)
    fmt.Print(" ")
    fmt.Printf("%02d", rm)
    fmt.Print(" ")
    fmt.Printf("%02d\n", rs)
    }
