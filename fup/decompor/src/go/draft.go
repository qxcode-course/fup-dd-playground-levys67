package main
import "fmt"
func main() {
    numero := 0
    fmt.Scan(&numero)
    numeros := make([]int,0)
    for numero > 0{
        digito := numero % 10
        numeros = append(numeros, digito )
        numero /= 10 
    }
    for i := len(numeros)-1; i >= 0; i--{
        if i != len(numeros) -1{
            fmt.Print(" ")
        }
        fmt.Print(numeros[i])
        if i == 0 {
            fmt.Println("")
        }
    } 

}