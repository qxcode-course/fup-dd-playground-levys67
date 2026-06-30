package main

import (
	"fmt"
)

func calcularcodigo (letra1, letra2, sinal string) string {
    numeros := map[string]int {
        "a" : 0,
        "b" : 1,
        "c" : 2,
        "d" : 3,
        "e" : 4,
        "f" : 5,
        "g" : 6,
        "h" : 7,
        "i" : 8,
        "j" : 9,
        "k" : 10,
        "l" : 11,
        "m" : 12,
        "n" : 13,
        "o" : 14,
        "p" : 15,
        "q" : 16,
        "r" : 17,
        "s" : 18,
        "t" : 19,
        "u" : 20,
        "v" : 21,
        "w" : 22,
        "x" : 23,
        "y" : 24,
        "z" : 25,
    }
    negativos := map[int]string {
        -1 : "z",
    }


    resultado := calculo(numeros[letra1],numeros[letra2], sinal)
    if resultado > 25 {
        resultado = resultado - 26
    }
    if resultado < 0 {
        for chave, valor := range negativos{
            if chave == resultado {
                return valor
            }
        }
    }
    for chave, valor := range numeros{
        if valor == resultado {
            return chave
        }
    }
    return ""
}

func calculo (valor1,valor2 int, sinal string) int {
    if sinal == "+" {
        return valor1 + valor2
    } else {
        return valor1 - valor2
    }
}


func main() {
    letra := ""
    sinal := ""
    letra2 := ""
    fmt.Scan(&letra,&sinal,&letra2)
    conta := calcularcodigo(letra,letra2,sinal)
    fmt.Println(conta)
}