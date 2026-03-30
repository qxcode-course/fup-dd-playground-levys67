package main
import "fmt"
func main() {
    var time int = 0
    fmt.Scan(&time)

    var hora int = time / 3600
    var min_seg int = (time % 3600) 
    var min int = min_seg / 60
    var sec int = min_seg % 60
    fmt.Printf("%d:%d:%d\n",hora,min,sec )
}
