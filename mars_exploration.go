package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var helpStr string;
    fmt.Scan(&helpStr)
    
    sos :=  "SOS"
    
    count := 0 //len(helpStr)/3
    for i := 0; i < len(helpStr); i+=3 {
        
        temp := helpStr[i:i+3]
        if temp != sos {
            
            for j := 0; j < 3; j++ {
                if temp[j] != sos[j] {
                    count++
                }
            }
        }
    } 
    
    fmt.Println(count)

}