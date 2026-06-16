package main
import "fmt"



func main() {
   cap := 0
   mov := 0
   fmt.Scan(&cap)
   passageiros := 0
   for {
    fmt.Scan(&mov)
    passageiros += mov
    if passageiros == 0 {
        fmt.Println("vazio")
    } else if passageiros > 0 && passageiros < cap {
        fmt.Println("ainda cabe")
    } else if passageiros >= cap && passageiros < cap*2{
        fmt.Println("lotado")
    } else if passageiros >= cap*2{
        fmt.Println("hora de partir")
        break
    }
    }
}