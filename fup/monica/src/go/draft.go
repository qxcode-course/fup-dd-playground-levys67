package main
import "fmt"
func main() {
    var idademae int 
    var idade1 int
    var idade2 int
    fmt.Scan(&idademae, &idade1, &idade2)
    total := (idade1 + idade2) 
    idade3 := idademae - total
    //fmt.Println(idade3)
    if idade3 > idade1 && idade3 > idade2 {
        fmt.Println(idade3)
    } else if idade3 < idade2 && idade2 > idade1 {
        fmt.Println(idade2)
    } else {
        fmt.Println(idade1)
    }

}
