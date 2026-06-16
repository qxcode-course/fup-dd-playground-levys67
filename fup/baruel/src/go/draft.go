package main
import "fmt"

func criarepreenchervetor(n int) []int {
    lista := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&lista[i])
    }
    return lista
}

func existe (lista [] int, valor int) bool {
    //contrado := false
    for i := 0; i < len(lista); i++{
        if lista[i] == valor {
            return true
        }
    }
    return false
}
func preencheralbum(lista [] int)( [] int ,[] int){
    album:= make([]int,0)
    repet:= make([]int,0)
    for i := 0; i < len(lista); i++{
        if !existe(album,lista[i]){
            album = append(album,lista[i])
        } else if existe(album,lista[i]){
            repet = append(repet, lista[i])
        }
    }
    return album, repet
}

func main() {
    figdoalbum := 0
    figtotal := 0
    fmt.Scan(&figdoalbum,&figtotal)
    figurinhas := criarepreenchervetor(figtotal)
   //mt.Println(figurinhas)
    album, repetidos:= preencheralbum(figurinhas)
    fmt.Print("[")
    for i := 0; i < len(repetidos) ; i++{
        fmt.Printf(" %d",repetidos[i])
    }
    fmt.Println(" ]")
    faltam := make([]int, 0)
    for i := 1; i <= figdoalbum; i++{
        if !existe(album,i){
            faltam = append(faltam, i)
        }
    }
    fmt.Print("[")
    for i := 0; i < len(faltam); i++{
        fmt.Printf(" %d", faltam[i])
    }
    fmt.Println(" ]")

}
