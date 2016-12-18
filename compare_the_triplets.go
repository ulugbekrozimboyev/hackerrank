package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    alice := make([]int,0, 3) 
    bob := make([]int,0, 3)
    
    alice_score := 0
    bob_score := 0
    
    
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            var tmp int
            fmt.Scan(&tmp)
            
            switch i {
                case 0 : alice = append(alice, tmp)
                default : bob = append(bob, tmp)
            }
        } 
    } 
    
    for j := 0; j < 3; j++ {
        
        if alice[j] > bob[j] {
            alice_score++
        } else if alice[j] < bob[j] {
            bob_score++
        }
    }
    
    fmt.Println(alice_score,bob_score)
}