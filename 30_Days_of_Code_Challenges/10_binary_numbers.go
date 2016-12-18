package main
import "fmt"
import "strconv"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n int64
    l := 0 
    fmt.Scan(&n)
    str := string(strconv.FormatInt(n, 2))
    //fmt.Println(str)
    for i := 0; i < len(str); i++ {
        tmp_l := 0
        //fmt.Println(str[i])
        if str[i] == 49 {
            for j := i; j < len(str); j++ {
                if str[j] == 49 {
                    tmp_l++
                } else {
                    break
                }
            } 
        }
        //fmt.Println(l,tmp_l)
        if l < tmp_l {
            l = tmp_l
        }
    }
    fmt.Println(l)
}