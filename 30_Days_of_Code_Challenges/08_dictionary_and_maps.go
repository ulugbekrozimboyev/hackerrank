package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

var m map[string]int
	
func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    m = make(map[string]int)
    
    var n int
    fmt.Scan(&n)
    
      
    
    for i := 0; i < n; i++ {
        var tmp int
        var s_tmp string
        
        fmt.Scan(&s_tmp)
        fmt.Scan(&tmp)
        
        m[s_tmp] = tmp
    }
    
    reader := bufio.NewReader(os.Stdin)
    for i := 0; i < n; i++ {
        s_tmp, _ := reader.ReadString('\n')
        //fmt.Println(s_tmp)
        if s_tmp[len(s_tmp)-1] == '\n'{
            s_tmp = s_tmp[:len(s_tmp)-1]   
        } 
        v, ok :=  m[strings.Trim(s_tmp," \n")]
        if ok {
            fmt.Printf("%s=%d\n",s_tmp, v )
        } else {
            fmt.Println("Not found")
        } 
    }
    
}