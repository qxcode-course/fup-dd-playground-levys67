 package main
import "fmt"

func calcularmmc (vetor []int) int{
    val1 := vetor[0]
    val2 := vetor [1]
    mmc := 0
    vezes := 2
    for {
        if val1 < val2 {
            val1 = (val1 * vezes)
        } else if val2 < val1 {
            val2 = (val2 * vezes)
        }
        if val1 != val2 {
            vezes ++
            val1 = vetor[0]
            val2 = vetor[1]
            continue
        } else {
            mmc = val1
            return mmc
        }
    }
}



func main() {
    numeros := make([]int, 2)
    for i := 0; i < len(numeros); i++{
        fmt.Scan(&numeros[i])
    }
    mmc := calcularmmc(numeros)
    fmt.Println(mmc)
}