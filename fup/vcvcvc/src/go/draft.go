package main

import (
	"bufio"
	"fmt"
	"os"
)

func separarletras (texto string) string {
    cv := ""
    for i := 0; i < len(texto); i++ {
        switch texto[i] {
        case 'a','e','i','o','u','A','E','I','O','U':
            cv += "v"
        case ' ':
            cv += " "
        default :
            cv += "c"
        }
    }
    return cv
}



func main() {
    texto := ""
    Scanner := bufio.NewScanner(os.Stdin)
    Scanner.Scan()
    texto = Scanner.Text()
    letras := separarletras(texto)
    fmt.Println(letras)
}