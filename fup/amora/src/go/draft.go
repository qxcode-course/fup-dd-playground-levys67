package main

import (
	"bufio"
	"fmt"
	"os"
)

func contarsub(texto, substring string) int {
    cont := 0
    for i := 0; i <= len(texto) - len(substring); i++{
        igual := true
        for j := 0; j < len (substring); i++{
            if texto[i+j] != substring[j] {
                igual = false
                break
            }
        }
        if igual {
            cont ++
        }
    }
    return cont
}



func main() {
    texto := ""
    sub := ""
    Scanner := bufio.NewScanner(os.Stdin)
    Scanner.Scan()
    texto = Scanner.Text()
    Scanner.Scan()
    sub = Scanner.Text()
    fmt.Println(texto,sub)
}