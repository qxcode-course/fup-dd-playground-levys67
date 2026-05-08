package main

import (
	"fmt"
	"unicode"
)
func main() {
    car := ""
    fmt.Scanln(&car)
    res := invert(car)
    fmt.Println(res)
    
}
func invert(car string) string {
    runes := []rune(car)[0]
    if unicode.IsLower(runes){
    return string(unicode.ToUpper(runes))
    }
    if unicode.IsUpper(runes) {
    return string(unicode.ToLower(runes))
    }
    return string(runes)
}