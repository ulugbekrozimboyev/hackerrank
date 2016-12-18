package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    var i uint32 = 4
    var d float32 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewReader(os.Stdin)


    // Declare second integer, double, and String variables.
    var my_i uint32  
    var my_d float32  
          
    // Read and save an integer, double, and String to your variables.
    //var my_s string
    fmt.Scan(&my_i)
    fmt.Scan(&my_d)
    
    my_s, _ := scanner.ReadString('\n')

    // Print the sum of both integer variables on a new line.
    fmt.Println(i+my_i)
    // Print the sum of the double variables on a new line.
    fmt.Printf("%.1f\n", d+my_d)
    // Concatenate and print the String variables on a new line
    // The 's' variable above should be printed first.    
    fmt.Println(s+my_s)
    
}