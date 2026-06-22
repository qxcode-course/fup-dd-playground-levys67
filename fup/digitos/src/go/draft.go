package main
import "fmt"
func main() {
    num := 0
    contato := 0
    fmt.Scan(&num, &contato)
    cont_texto := fmt.Sprint(contato)
    num_txt := fmt.Sprint(num)
    qtd := 0
    for i := 0; i < len(cont_texto); i++{
        if cont_texto[i] == num_txt[0] {
            qtd ++
        }
    }
    fmt.Println(qtd)
}