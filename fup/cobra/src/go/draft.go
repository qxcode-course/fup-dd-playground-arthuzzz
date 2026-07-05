package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    var x, y int
    fmt.Scan(&x, &y)

    var c string
    fmt.Scan(&c)

    var s int
    fmt.Scan(&s)

    if c == "R" {
        x = x + s
        if x >= n {
            x = x % n
        }
    } else if c == "L" {
        x = x - s
        if x < 0 {
            x = (x%n + n) % n
        }
    } else if c == "D" {
        y = y + s
        if y >= n {
            y = y % n
        }
    } else if c == "U" {
        y = y - s
        if y < 0 {
            y = (y % n + n) % n
        }
    }

    fmt.Println(x, y)


}