package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n int
    fmt.Scan(&n)
    for i := 0 ; i < n; i++ {
        for j := 0 ; j < n; j++ {
            if i+j>=n-1 { 
                fmt.Print("#") 
            } else {
                fmt.Print(" ") 
            }
        }
        fmt.Println()
    }
}