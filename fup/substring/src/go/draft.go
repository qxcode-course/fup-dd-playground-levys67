package main

import (
	"bufio"
	"fmt"
	"os"
    "strconv"
)
func main() {
    var text string
    ind := 0
    qtd := 0
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text = scanner.Text()
    indt := scanner.Text()
    ind, _ = strconv.Atoi(indt)

    listtexto := []rune(text)
    //caracteres := 0
    for i := 0; i < len(text); i++{
        fmt.Printf("%c", listtexto[i])
    }
    fmt.Println()
    for i := 0; i < qtd; i++{
        fmt.Printf("%c", listtexto[i])
    }
    fmt.Println(ind)
    fmt.Println(qtd)
}