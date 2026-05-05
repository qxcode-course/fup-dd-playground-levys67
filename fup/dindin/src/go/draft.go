package main

import ("fmt")
func main() {
    qtde := 0
    fmt.Scan(&qtde)
    choc := 0
    lim := 0
    man := 0
    tar := 0
    sabor := ""
    turno := ""
    for i := 0; i < qtde; i++ {
        fmt.Scan(&sabor, &turno)
        if sabor == "c" {
            choc += 1
        } else {
            lim += 1
        }
        if turno == "m" {
            man += 1
        } else {
            tar += 1
        }
    }
    if choc > lim {
        fmt.Println("c")
    } else if lim > choc {
        fmt.Println("l")
    } else {
        fmt.Println("empate")
    }
    if man > tar {
        fmt.Println("t")
    } else if tar > man {
        fmt.Println("m")        
    } else {
        fmt.Println("empate")
    }
}
    
  
    