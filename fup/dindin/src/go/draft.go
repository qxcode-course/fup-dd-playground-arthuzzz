package main
import "fmt"
func main() {
    var dd int
    //var c, l int
    //var m, t int

    fmt.Scan(&dd)

    choc:= 0
    lima:= 0
    manha:= 0
    tarde:= 0


    sabor:= make([]string, dd)
    turno:= make([]string, dd)

    for i := 0; i < dd; i++ {
    fmt.Scan(&sabor[i], &turno[i])
    }

    for i:= 0; i < dd; i++ {
        if sabor[i] == "c" {
            choc++
        }
        if sabor[i] == "l" {
            lima++
        }
        if turno[i] == "m" {
            manha++
        }
        if turno[i] == "t" {
            tarde++
        }

    }

    if choc > lima {
        fmt.Println("c")
    } else if choc < lima {
        fmt.Println("l")
    } else if choc == lima {
        fmt.Println("empate")
    }
    if manha < tarde {
        fmt.Println("m")
    } else if manha > tarde {
        fmt.Println("t")
    } else if manha == tarde {
        fmt.Println("empate")
    }
    
}