package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    var nums []int = make([]int, N)

   for i := 0; i < N; i++ {
		fmt.Scan(&nums[i])
    }

    fmt.Print("[")

    for i := 0; i < N; i++ {
        if i > 0 {
            fmt.Print(", ")
        }
        fmt.Print(nums[i])
    }
    fmt.Print("]\n")
}


