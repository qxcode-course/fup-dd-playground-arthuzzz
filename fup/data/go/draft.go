package main
import "fmt"
func main() {
    var hh, mm, dd, MM, aa int
    fmt.Scan(&hh, &mm, &dd, &MM, &aa)
    
    ano := aa % 100

    fmt.Printf("%02d:%02d %02d/%02d/%02d\n", hh, mm, dd, MM, ano)
}
