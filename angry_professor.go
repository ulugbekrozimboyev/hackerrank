package main
import "fmt"

func main() {
 	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var countClass int;
	fmt.Scan(&countClass)
	// var answers = make([]string, 0, 10)

	for i := 0; i < countClass; i++ {
		var studensCount, studentsMinCount int
		var studentsThatArenotLate int

		fmt.Scan(&studensCount, &studentsMinCount)
		for j := 0; j < studensCount; j++ {
			var studentLateTime int
			fmt.Scan(&studentLateTime)

			if studentLateTime <= 0 {
				studentsThatArenotLate++
			}
		}
		// fmt.Printf("studentsThatArenotLate %d\n",studentsThatArenotLate)
		if studentsThatArenotLate >= studentsMinCount{
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}

	} 
}