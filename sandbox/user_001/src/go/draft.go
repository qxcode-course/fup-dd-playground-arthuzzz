func mostrar_vetor(arr []int, sep string) {
    fmt.Print("[")
    for i, valor:= range arr {
        if i != 0 {
            fmt.Print(sep)
        }
        fmt.Printf("%v", valor)
    }
    fmt.Print("]\n")
}