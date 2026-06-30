package main

import (
	"bufio"
	"fmt"
	"os"
)

func contarsub(texto, substring string) int {
    cont := 0
    runastexto := []rune(texto)
    runassub := []rune(substring)
    for i := 0; i <= len(runastexto) - len(runassub); i++{
        igual := true
        for j := 0; j < len (runassub); j++{
            if runastexto[i+j] != runassub[j] {
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
    contagem := contarsub(texto,sub)
    fmt.Println(contagem)
}