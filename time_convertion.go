package main
import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var timeStr string
    fmt.Scan(&timeStr)

    timeType := timeStr[8:]
    timeStr = timeStr[:8]
 	str := strings.Split(timeStr,":")

 	var resTime string
 	switch {
 		case (timeType == "AM" && timeStr == "12:00:00") : 
 			resTime = "00:00:00";
 		case (timeType == "AM" && str[0] == "12") : 
 			resTime = fmt.Sprintf("00:%s:%s", str[1], str[2]);
 		case timeType == "AM" : 
 			resTime = timeStr;
 		case timeType == "PM" && str[0] == "12" : 
 			resTime = timeStr
 		case timeType == "PM" : 
	 		hour, err := strconv.Atoi(str[0])
	 		if err != nil {
	 			fmt.Print(err)
	 		}
	 		resTime = fmt.Sprintf("%d:%s:%s", hour + 12, str[1], str[2]);
 	}   
 	fmt.Println(resTime)
}