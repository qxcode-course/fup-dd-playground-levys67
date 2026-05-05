package main
import "fmt"
func main() {
    dindin := 0
    sabor, turno := "", ""
    fmt.Scan(&dindin, &sabor, &turno)
    lista := make([]string, dindin)
    for i := 0; i < dindin; i++ {
        fmt.Scan(&lista[i])
    }
    
    chocolate := 0
    limão := 0
    manha:= 0
    tarde := 0
    for _, v := range lista{
        if v == "c" {
            chocolate += 1
        } else if v == "l" {
            limão += 1
        } else if v == "m" {
            manha += 1
        }else if v == "t" {
            tarde += 1
        }
    }

    if chocolate > limão {
        fmt.Println("c")
    } else if limão > chocolate {
        fmt.Println("l")
    }else if chocolate == limão {
        fmt.Println("empate")
    }
    if manha < tarde {
        fmt.Println("m")
    } else if tarde < manha {
        fmt.Println("t")
    } else if manha == tarde {
        fmt.Println("empate")
    }
}