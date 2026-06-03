package main
import "fmt"
func main() {
    a := 0
    b := 0
    c := 0
    h := 0
    l := 0
    fmt.Scan(&a, &b, &c, &h, &l)
    face1 := a * b
    face2 := a * c 
    face3 := b * c
    janela := h * l
    if janela >= face1 || janela >= face2 || janela >= face3 {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}
