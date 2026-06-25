package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func invertercase(letras []rune) string {
    inverso := make([]rune, 0)
    for i := 0; i < len(letras); i++{
        if unicode.IsLower(letras[i]) {
            inverso = append(inverso, unicode.ToUpper(letras[i]))
        } else {
            inverso = append(inverso, unicode.ToLower(letras[i]))
        }
    }
    return string(inverso)
}



func main() {
    texto := ""
    Scanner := bufio.NewScanner(os.Stdin)
    Scanner.Scan()
    texto = Scanner.Text()
    letras := []rune(texto)
    inverso := invertercase(letras)
    fmt.Println(inverso)
}