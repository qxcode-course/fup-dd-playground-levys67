package main
import "fmt"

func criarvetor (num int) []int{
    vetor := make([]int,  num)
    for i := 0; i < num; i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}


func main() {
    num := 0
    fmt.Scan(&num)
    vetor := criarvetor(num)
    fmt.Print("[")
    for i, v := range vetor {
        if i != 0 {
            fmt.Print(", ")
        }
        if v == 1 {
            fmt.Print("A")
        } else if v == 11 {
            fmt.Print("J")
        }else if v == 12{
            fmt.Print("Q")
        }else if v == 13{
            fmt.Print("K")
        } else {
            fmt.Print(v)
        }
    }
    fmt.Println("]")

}