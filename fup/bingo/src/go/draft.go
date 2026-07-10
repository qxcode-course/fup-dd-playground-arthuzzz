package main
import "fmt"
func main() {
    cartela := [4][4] int {
        {1, 9, 27, 23},
        {34, 20, 37, 47}, 
        {30, 87, 55, 69},
        {13, 60, 99, 66},
    }
    var nums [6]int 

    for i:= 0; i < 6; i++ {
        fmt.Scan(&nums[i])
    }
    c:= 0

    for i:= 0; i < 6; i++ {
        for a := 0; a < 4; a++ {
            for b := 0; b < 4; b++ {
            if nums[i] == cartela[a][b] {
                c++
            }
        }
        }
    }
    fmt.Println(c)
}