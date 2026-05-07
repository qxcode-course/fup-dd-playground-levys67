package main
import "fmt"
func main() {
    mais_novo := 0
    qtd := 0
    fmt.Scan(&mais_novo, &qtd)
    idades := ((qtd - 1) * 2) + mais_novo
    if idades % 2 == 0 {
        for i := mais_novo; i <= idades; i++ {
            if i%2 == 0 {
                fmt.Println(i)
            }
        }
    } else if idades % 2 != 0{
    for i := mais_novo; i <= idades; i++ {
        if i%2 != 0 {
            fmt.Println(i)
        }
    }
}
}
