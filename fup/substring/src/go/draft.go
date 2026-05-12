package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
    var text string
    ind := 0
    qtd := 0
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text = scanner.Text()
    fmt.Scan(&ind, &qtd)
    listtexto := []rune(text)
    caracteres := 0
    for i := 0; i <= len(text)-1; i++{
        if i >= ind {
            fmt.Printf("%c", listtexto[i])
        }
        caracteres += 1
        if caracteres > qtd{
            break
        }
    }
    fmt.Println()
}