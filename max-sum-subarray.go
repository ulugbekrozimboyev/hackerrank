package main

import "fmt"

func checkBigger(a,b int) int {
    if a > b {
         return a
    } else {
        return b
    }        
}

func getMaxArrSum(arr []int) int {
    
    var result int
    for i := 0; i < len(arr); i++ {
        if arr[i] == 0 {
            continue
        }
        var hasNegative bool
        var temp int;
        for j := i; j < len; j++ {
            if arr[j] == 0 {
                result = checkBigger(result, temp)
                temp = 0
                continue
            } else {
                temp += arr[j]
                result = checkBigger(result,temp)
            }
            if arr[j] < 0 {
                hasNegative = true
            }
        }
        if !hasNegative {
            break
        }
    }

    return result

}

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var len int
    fmt.Scan(&len)
    arr := make([]int, 0, 100000)
    var result int
    var lastIndex int
    for i := 0; i < len; i++ {
        var a int
        fmt.Scan(&a)
        if a == 0 {
            temp := getMaxArrSum(arr[lastIndex : ])
            result = checkBigger(result,temp)
        }
        arr = append(arr, a)

        result = getMaxArrSum(arr)
    }
    fmt.Println(result)
}