package main
import "fmt"
func main() {
    
    var total int

    fmt.Scan(&total)

    horas := total / 3600
    resto := total % 3600

    minutos := resto / 60
    segundos := resto % 60

    fmt.Printf("%d:%d:%d\n", horas, minutos, segundos)
}
