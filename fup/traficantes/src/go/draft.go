package main

import (
	"bufio"
	"fmt"
	"os"
)

func substituir(texto, alvo, subs string) string {
    var textoalt []byte//armazena letras
    i := 0//indice que começa
    for i < len(texto) {//condicional externa e percorre texto
        encontrou := true//variavel de controle
        if i+len(alvo) <= len(texto) {//verifica se cabe dentro da string
            for j := 0; j < len(alvo); j++ {//percorre alvo para substituir
                if texto[i+j] != alvo[j] {//verifica se os caracteres batem
                    encontrou = false
                    break
                }
            }
        } else {
            encontrou = false
        }
        if encontrou {
            for j := 0; j < len(subs); j++ {//percorre substring que ira substituir
                textoalt = append(textoalt, subs[j])//adiciona letras
            }
            i += len(alvo)//cobre o espaço da palavra antiga 
        } else {
        textoalt = append(textoalt, texto[i])//adiciona elementos do texto base
        i++// finaliza o loop adicionando +1 ao indice
    }
}
    return string(textoalt)
}

func main() {
    texto := ""
    alvo := ""
    subs := ""
    Scanner := bufio.NewScanner(os.Stdin)
    Scanner.Scan()
    texto = Scanner.Text()
    Scanner.Scan()
    alvo = Scanner.Text()
    Scanner.Scan()
    subs = Scanner.Text()
    frase := substituir(texto,alvo,subs)
    fmt.Println(frase)
}