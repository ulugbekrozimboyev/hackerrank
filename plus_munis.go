package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n int
    fmt.Scan(&n)
    var pos,neg,zeros float32 
    
    for i := 0 ; i < n; i++ {
        var d int
        fmt.Scan(&d)
        switch {
            case d > 0: pos++;
            case d == 0: zeros++;
            case d < 0: neg++;
        }
    }

    //var pf float32 = float32(pos/n)

    fmt.Printf("%.6f\n", pos/float32(n))
    fmt.Printf("%.6f\n", neg/float32(n))
    fmt.Printf("%.6f\n", zeros/float32(n))
}