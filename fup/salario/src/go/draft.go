package main
import "fmt"
func main() {
    salario := 0
    fmt.Scan(&salario)
    aumento1 := (salario * 20) / 100
    aumento2 := (salario * 15) / 100
    if salario <= 1000 {
        fmt.Printf("%.2f\n",float64(salario + aumento1))
    } else if salario > 1000 && salario <= 1500 {
        fmt.Printf("%.2f\n",float32(salario + aumento2))
    }
}