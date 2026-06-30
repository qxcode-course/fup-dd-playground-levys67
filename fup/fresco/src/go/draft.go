package main

import (
	"bufio"
	"fmt"
	"os"
)

func juntarvogais(letras string) string {
runas := []rune (letras)
juntas := ""
for i := 1; i < len(runas)-1; i++ {
    if runas[i] == ' ' && (runas[i-1] == 'a' || runas[i-1] == 'e' || runas[i-1] == 'i' && runas[i-1] == 'o' || runas[i-1] == 'u') && runas[i+1] == runas[i-1]{ 
        continue
    } 
    juntas += string(runas[i])    
}
return juntas
}
    


func main() {
    texto := ""
    Scanner := bufio.NewScanner(os.Stdin)
    Scanner.Scan()
    texto = Scanner.Text()
    juntas := juntarvogais(texto)
    fmt.Println(juntas)
}