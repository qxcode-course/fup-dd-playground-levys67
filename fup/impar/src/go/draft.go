package main
import "fmt"
func main() {
    par_ou_impar := 0
    dedos1 := 0
    dedos2 := 0
    fmt.Scan(&par_ou_impar,&dedos1,&dedos2)
    if par_ou_impar == 0 && (dedos1 + dedos2) % 2 == 0 {
        fmt.Println(0)
    } else if par_ou_impar == 0 && (dedos1 + dedos2) % 2 != 0{
        fmt.Println(1)
    }else if par_ou_impar == 1 && (dedos1 + dedos2) % 2 == 0{
        fmt.Println(1)
    }else if par_ou_impar == 1 && (dedos1 + dedos2) % 2 != 0{
        fmt.Println(0)
    }
}
