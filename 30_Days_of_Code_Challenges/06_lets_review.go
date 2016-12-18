package main
import (
    "fmt"
    "os"
    "bufio"
)
func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n int
    fmt.Scan(&n)
    
    reader := bufio.NewReader(os.Stdin)
    
    for i:= 0; i < n; i++ {
        var toq, juft string
        str, _ := reader.ReadString('\n')
        for j := 0; j < len(str); j++ {
            if str[j] != '\n' {
                if j % 2 == 0 {
                    juft +=string(str[j])                    

                } else {
                    toq += string(str[j])
                } 
            }           
        }
        fmt.Println(juft, toq)
    } 
}