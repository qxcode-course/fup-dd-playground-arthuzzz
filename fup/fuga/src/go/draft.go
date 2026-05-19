package main

import "fmt"

func main() {
    var H, P, F, D int
    fmt.Scan(&H, &P, &F, &D)

    if D == -1 {
        if F > H {
            if P < F && P > H {
                fmt.Println("N")
            } else {
                fmt.Println("S")
            }

        } else {
            if P < F || P > H {
                fmt.Println("N")
            } else {
                fmt.Println("S")
            }
        }

    } else {
        if F < H {

            if P > F && P < H {
                fmt.Println("N")
            } else {
                fmt.Println("S")
            }

        } else {
            if P > F || P < H {
                fmt.Println("N")
            } else {
                fmt.Println("S")
            }
        }
    }
}