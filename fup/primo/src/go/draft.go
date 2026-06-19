package main

import (
	"fmt"
	"math"
)
 func Primo(num int) int {

     sqrt := int(math.Sqrt(float64(num)))

     if num < 2 {
        return 0
     }

for i:= 2; i <= sqrt; i++ {
        if num%i == 0 {
            return 0
        }
        }
           return 1
    }

func main() {
    var num int
    fmt.Scan(&num)
    
    fmt.Println(Primo(num))
}