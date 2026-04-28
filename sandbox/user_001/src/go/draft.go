package main
import "fmt"
func filtrar_impares(nums []int) []int{
	impares := make([]int,0 ,len(nums))
	for _, valor := range nums {
		if valor % 2 == 1{
			impares := append(impares, valor)
		}
	}
	return impares
}
func index (nums []int,valor int)int {
    for i, elem := range nums {
		if elem == valor {
			return i
		}
    }
	return -1
}
func contain (nums []int, valor int)bool {
	for _, elem := range nums {
		if elem == valor {
			return true
		}
	}
	return false
}
func count (nums []int, valor int) int{
    contador := 0
    for _, elem := range nums {
	    if elem == valor {
		    contador += 1
	    }
    }
return contador
}
func separar_figurinhas (montante[]int) ([]int, []int)

func main() {
    
}

