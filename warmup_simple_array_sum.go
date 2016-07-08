package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var size int
    //var arr = make([]int, 0, 10)
    var summ int64 = 0
    fmt.Scan(&size)
    for i := 0; i < size; i++ {
    	var d int
        fmt.Scan(&d)
        summ += int64(d)
    }
    fmt.Println(summ)
}