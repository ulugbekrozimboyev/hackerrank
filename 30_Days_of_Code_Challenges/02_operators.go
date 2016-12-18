package main
import "fmt"
import "math"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var mealCost float64
    var tipPercent float64
    var taxPercent float64
    fmt.Scan(&mealCost)
    fmt.Scan(&tipPercent)
    fmt.Scan(&taxPercent)
    
    totalPrice := mealCost *(100+tipPercent+taxPercent)/100
    
    //fmt.Println(mealCost)
    //fmt.Println(tipPercent)
    //fmt.Println(taxPercent)
    //fmt.Println(totalPrice - math.Floor(totalPrice))
    //fmt.Println(math.Ceil(totalPrice) - totalPrice)
    
    if totalPrice - math.Floor(totalPrice) < math.Ceil(totalPrice) - totalPrice {
        fmt.Printf("The total meal cost is %d dollars.", int(totalPrice))
    } else {
        fmt.Printf("The total meal cost is %d dollars.", int(math.Ceil(totalPrice)))
    }
    
    
    //fmt.Printf("The total meal cost is %f dollars.",  totalPrice)
}