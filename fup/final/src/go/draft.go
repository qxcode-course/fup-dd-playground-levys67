package main
import "fmt"
func main() {
    nota1 := 0
    nota2 := 0
    notafinal := 0
    fmt.Scan(&nota1, &nota2, &notafinal)
    media := float64((nota1 + nota2))/2
    mediafinal := float64((media + float64(notafinal))/2)
    if media >= 7 {
        fmt.Println("aprovado")
    } else if media < 4 {
        fmt.Println("reprovado")
    } else  {
        if  mediafinal >= 5{
            fmt.Println("aprovado na final")
        } else{
            fmt.Println("reprovado na final")
        }
    }
}
