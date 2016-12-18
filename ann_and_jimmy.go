package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n int
    fmt.Scan(&n)
    max := 0
    for i := 1; i <= n; i ++ {
        for j := 1; j <=n ; j++ {
            for k := 1; k <= n; k++ {
                if i + j + k == n && max < i*j*k {
                    max = i*j*k
                }
            }
        }
    }
    
    fmt.Println(max)
}