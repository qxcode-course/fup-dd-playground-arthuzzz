package main
import "fmt"
    type Fator struct {
        num int
        qtd int
    }
func calculo(num int) []Fator {
    var fatores []Fator
    for i := 2; i <= num; i++ {
        qtd := 0

        for num % i == 0 {
            qtd++
            num = num / i
        }
        if qtd > 0 {
            fatores = append(fatores, Fator {
                num: i,
                qtd: qtd, 
            })
        }
    }
        return fatores
     }



func main() {

    var n int
    fmt.Scan(&n)

    x := calculo(n)

    for _, f := range x {
        fmt.Println(f.num, f.qtd)
    }
}