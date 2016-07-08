package main
import "fmt"
import "math"

func main() {
    var n int
    var s1 int64 = 0
    var s2 int64 = 0
    
    fmt.Scan(&n)
    // fmt.Println(n)
    var a = make([][]int64, 10)
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    for i := 0; i < n; i++ {
        a[i] = make([]int64, n)
        for j :=0; j < n; j++ {
            fmt.Scan(&a[i][j])
        }
        s1 += a[i][i]
        s2 += a[i][n-i-1]
    }
    
    
    
    fmt.Println(math.Abs(float64(s1-s2)))
}