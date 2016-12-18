package main
import "fmt"

func insertionSort(arr *[]int, n int){
    newArr := * arr
    for i := 0 ; i < n; i++ {
        tmp := newArr[i]
        j := i
        
        for ;  j > 0 && tmp < newArr[j-1] ; {
            newArr[j] = newArr[j-1]
            j -= 1
            str := fmt.Sprint(newArr)
            fmt.Println(str[1:len(str)-1])

        }
        newArr[j] = tmp
    } 
    
    *arr = newArr
}

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n int
    fmt.Scan(&n)
    arr := make([]int, 0, n)
    
    for i := 0; i < n; i++ {
        var tmp int
        fmt.Scan(&tmp)
        arr = append(arr, tmp)
    } 
    
    insertionSort(&arr, n)
    str := fmt.Sprint(arr)
    fmt.Println(str[1:len(str)-1])
}