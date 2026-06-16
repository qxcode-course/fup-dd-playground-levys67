package main
import "fmt"
func main() {
    salario := 0.0
    fmt.Scan(&salario)
    aumento1 := (salario * 20) / 100
    aumento2 := (salario * 15) / 100
    aumento3 := (salario * 10) / 100
    aumento4 := (salario * 5) / 100
    if salario <= 1000 {
        fmt.Printf("%.2f\n",salario + aumento1)
    } else if salario > 1000 && salario <= 1500 {
        fmt.Printf("%.2f\n",salario + aumento2)
    } else if salario > 1500 && salario <= 2000 {
        fmt.Printf("%.2f\n",salario + aumento3)
    } else{
        fmt.Printf("%.2f\n",salario + aumento4)
    } 

    
}