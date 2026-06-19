package main
import "fmt"
func main() {
    numero := 0
    fmt.Scan(&numero)
    encontrado := false
    for i := 0; i <= numero; i++{
        if numero == i*i && i*i != 0 {
            fmt.Println("sim")
            encontrado = true
        } 
        
        
    }
    if !encontrado {
        fmt.Println("nao")
    }
}