package main

import (
	"bufio"
	"fmt"
	"os"
)

func separarletras (letras []rune) (string,string){
    vogais := ""
    consoantes := ""
    for i := 0; i < len(letras); i++ {
        if letras[i] == 'a' || letras[i] == 'e' || letras[i] == 'i' || letras[i] == 'u' || letras[i] == 'o' {
            vogais += string(letras[i])
        } else if letras[i] != ' ' {
            consoantes += string(letras[i])
        }
    }
    return vogais, consoantes
}




func main() {
    texto := ""
    Scanner := bufio.NewScanner(os.Stdin)
    Scanner.Scan()
    texto = Scanner.Text()
    letras := []rune(texto) 
    vogais, consoantes := separarletras(letras)
    fmt.Printf("%s\n", vogais)
    fmt.Printf("%s\n", consoantes) 
}