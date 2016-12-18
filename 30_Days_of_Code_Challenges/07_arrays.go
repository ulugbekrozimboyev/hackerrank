package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n int
    fmt.Scan(&n)
    
    arr := make([]int, 0, n)
    
    for i := 0; i<n; i++{
        var tmp int
        fmt.Scan(&tmp)
        arr = append(arr, tmp)
    }
    
    for i := n-1; i >=0 ; i-- {
        fmt.Printf("%d ",arr[i])
    }
    
}