package main

import (
	"fmt"
	"unicode"
)

func invertercaractere(caractere []rune) string {
    if unicode.IsLower(caractere[0]) {
        return string(unicode.ToUpper(caractere[0]))
    } else {
        return string(unicode.ToLower(caractere[0]))

    }
}


func main() {
    car := ""
    fmt.Scan(&car)
    runa := []rune(car)
    inv := invertercaractere(runa)
    fmt.Println(inv)
}
