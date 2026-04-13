package main
import "fmt"
func main() {
    mae := 0
    filho1 := 0
    filho2 := 0
    fmt.Scan(&mae,&filho1,&filho2)
    filho3 := mae -(filho1 + filho2)
    if filho3 > filho1 && filho3 > filho2{
        fmt.Println(filho3)
    } else if filho2 > filho3 && filho2 > filho1{
        fmt.Println(filho2)
    } else if filho1 > filho3 && filho1 > filho2{
        fmt.Println(filho1)
    }
}
