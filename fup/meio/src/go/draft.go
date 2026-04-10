package main
import "fmt"
func main() {
    var scratt1, scratt2,scratt3 int
    fmt.Scan(&scratt1, &scratt2, &scratt3)

    if (scratt2 > scratt1 && scratt1 > scratt3) || (scratt2 < scratt1 && scratt1 < scratt3) {
        fmt.Println(scratt1)
    } else if (scratt1 > scratt2 && scratt2 > scratt3) || (scratt1 < scratt2 && scratt2 < scratt3) {
        fmt.Println(scratt2)
    } else  {
        fmt.Println(scratt3)
    }
}
