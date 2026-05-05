a <= b {
        for i:= a; i <= b; i++ {
            if (i%2 == 0) && (i%3 == 0) {
                soma += i
                fmt.Println(soma)
            }
           
        }
    } else if a > b {
        fmt.Println("invalido")
    }





    a > b {
        fmt.Println("invalido")
        return
    }
    soma:= 0

    for i:= a; i <= b; i++ {
        if i % 2 == 0 && i % 3 == 0 {
            soma += i
        }

    }

    fmt.Println(soma)