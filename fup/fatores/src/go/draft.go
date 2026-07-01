package main
import "fmt"
func main() {
    num := 0
    fmt.Scan(&num)
    numero := num
    contagem := map[int]int {
        2 : 0,
        5 : 0,
        7 : 0,
        11 : 0,
        23 : 0,

    }

    for range num{
        for i := 2; i <= numero; i++{
            if numero % i == 0{
                numero = numero / i
                contagem[i] ++
                i --
            }
        }

        if numero == 1 {
            break
        }
    }
    for chave, valor := range contagem{
        if valor != 0 {
            fmt.Printf("%d %d\n",chave, valor)
        }
    }
}