package main
import "fmt"
func main() {
    com := 160
    alt := 70
    cor1 := 0
    cor2 := 0
    fmt.Scan(&cor1, &cor2)
    //original := com * alt
    base1 := com - cor1
    topo1:= com - cor2
    base2 := cor1
    topo2 := cor2
    area1 := (base1 + topo1) * alt / (2)
    area2 := (base2 + topo2) * alt / (2)
    if area1 > area2 {
        fmt.Println(2)
    } else if area2 > area1 {
        fmt.Println(1)
    } else {
        fmt.Println(0)
    }
}
