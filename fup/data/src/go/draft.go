package main
import "fmt"
func main() {
    hora := 0;minuto := 0;dia := 0;mes := 0;ano := 0
    fmt.Scan(&hora,&minuto,&dia,&mes,&ano)
    fmt.Printf("%02d:%02d %02d/%02d/%02d\n",hora,minuto,dia,mes,ano%100)
}
