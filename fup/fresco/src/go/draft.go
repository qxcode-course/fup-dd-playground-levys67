package main

import (
	"bufio"
	"fmt"
	"os"
)

func juntarvogais(letras string) string {
runas := []rune (letras)
juntas := ""//armazena letras
i := 0
    for i < len(runas){
        if i+1 < len(runas) && runas[i] == ' ' && ehvogal(runas[i+1]){//verifica se o elemento atual e um espaço e o proximo uma vogal 
            i += 2//pula 2 eleementos
            continue
        }
        juntas += string(runas[i])// realiza a concatenação
        i ++
    }
    return juntas
}
    
func ehvogal (c rune) bool {
    if c == 'a' || 
    c == 'e' ||
    c == 'i' ||
    c == 'o' ||
    c == 'u' {
        return true
    }
    return false
}



func main() {
    texto := ""
    Scanner := bufio.NewScanner(os.Stdin)
    Scanner.Scan()
    texto = Scanner.Text()
    juntas := juntarvogais(texto)
    fmt.Println(juntas)
}