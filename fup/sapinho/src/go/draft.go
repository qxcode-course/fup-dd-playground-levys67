package main
import "fmt"
func main() {
    profundidade := 0
    salto := 0
    escorregao := 0
    fmt.Scan(&profundidade, &salto, &escorregao)
    progresso := 0
    saltos := 0
    for { //pqp juro que fis muito sem querer,obs esse for aberto é muito roubado xD
        if saltos != 0 {
            progresso -= escorregao
        }
        fmt.Printf("%d ", progresso)
        progresso += salto
        saltos ++
        if progresso >= profundidade{
            fmt.Println("saiu")
            break
        }
        fmt.Printf("%d\n", progresso)



    }
}