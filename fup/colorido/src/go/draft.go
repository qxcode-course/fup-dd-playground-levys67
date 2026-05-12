package main

import (
	"crypto/elliptic"
	"fmt"
)
func main() {
    pedra := 0
    pe := ""
    fmt.Scan(&pedra, &pe)
    fmt.Print("[")
    //localdapedra := 0
    if pe == "d" {
        for i := 0; i <= 10; i++ {
            if i != pedra && i % 2 != 0 {
                fmt.Printf(" %de", i)
            } else if i != pedra && i % 2 == 0 && i != 10{
                fmt.Printf(" %dd", i)
            }
            if i == pedra {
                break
            }
        }
        for i := 0; i <= 10; i ++{
            if i > pedra && i % 2 != 0{
                fmt.Printf(" %dd", i)
            } else if i > pedra && i % 2 == 0 && i != 10{
                fmt.Printf(" %de", i)
            }
        }
        for i := 0; i <= 10; i++{
            if i == 10 {
                fmt.Print(" ceu")
            }
        }
    } else if pe == "e" {
        for i := 0; i <= 10; i++ {
            if i != pedra && i % 2 != 0 {
                //parei aqui
            }
        }
    }
    fmt.Println(" ]")
}
