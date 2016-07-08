package main
import (
    "fmt"
    "strings"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var helpStr string;
    fmt.Scan(&helpStr)
    
    /*count := len(helpStr)/3
    for i := 0; i < len(helpStr); i+=3 {
        temp := helpStr[i:i+3]
        fmt.Println(temp)
        if strings.ToUpper(temp) == "SOS" {
            count--
        }
    }*/ 

    /*arr := strings.Split(helpStr,"SOS")

    count:=0
    fmt.Printf("%q\n",arr)
    for i := 0; i < len (arr); i++ {
        if arr[i] != "" {
            count++
        }
    }*/

    //count := len(strings.Replace(helpStr,"SOS","", -1))/3

    count := len(strings.Join( strings.Split(helpStr,"SOS"), "") )/3

    fmt.Println(count)

}