package main
import "fmt"
func main() {
    nota1 := 0
    nota2 := 0
    notafinal := 0
    fmt.Scan(&nota1, &nota2, &notafinal)
    media := float64((nota1 + nota2))/2
    mediafinal := float64((nota1 + nota2 + notafinal))/3
    if media >= 7 {
        fmt.Println("aprovado")
    } else if media < 4 {
        fmt.Println("reprovado")
    } else if (media < 7 && media >= 4 && mediafinal >= 5) {
        fmt.Println("aprovado na final")
    } else if (media < 7 && media >= 4 && mediafinal < 4){
        fmt.Println("reprovado na final")
    }
}
