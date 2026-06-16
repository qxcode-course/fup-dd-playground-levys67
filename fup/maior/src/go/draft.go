package main
import "fmt"
func main() {
    chute := 0.0
    direçao := ""
    valor := 0.0
    fmt.Scan(&chute,&direçao,&valor)
    if direçao == "M" && chute > valor || direçao == "m" && chute < valor {
        fmt.Println("primeiro")
    } else if direçao == "M" && chute < valor || direçao == "m" && chute > valor {
        fmt.Println("segundo")
    }

}