package main
import "fmt"
import "math/rand"

func partition ( a * []int,start ,end int) int {
    A := *a
    i := start + 1;
    piv := A[start] ;            //make the first element as pivot element.
    for j :=start + 1; j <= end ; j++   {
    /*rearrange the array by putting elements which are less than pivot
       on one side and which are greater that on other. */

          if A[ j ] < piv {
                tmp := A[i]
			    A[i] = A[j]
			    A[j] = tmp
            //swap (A[ i ],A [ j ]);
            i += 1;
        }
   }
    tmp := A[i-1]
    A[i-1] = A[start]
    A[start] = tmp

    *a = A

   //swap ( A[ start ] ,A[ i-1 ] ) ;  //put the pivot element in its proper place.
   return i-1;                      //return the position of the pivot
}

func quick_sort (a *[]int ,start,end  int) {
	
   if start < end {
        //stores the position of pivot element
         piv_pos := partition (a,start , end ) ;     
         quick_sort (a,start , piv_pos -1);    //sorts the left side of pivot.
         quick_sort (a,piv_pos +1 , end) ; //sorts the right side of pivot.
   }

   
}

func rand_partition ( a *[]int , start , end int ) {
    //chooses position of pivot randomly by using rand() function .
    random := rand.Intn(end+1) //start + rand( )%(end-start +1 ) ;
    A := *a

    tmp := A[random]
    A[random] = A[start]
    A[start] = tmp
    //swap ( A[random] , A[start]) ;        //swap pivot with 1st element.
    partition(a,start ,end) ;       //call the above partition func.

    *a = A
}

func main() {
 	//Enter your code here. Read input from STDIN. Print output to STDOUT
	
	var n int
	n = 7
	a := []int{5,8,1,3,7,9,2}

	rand_partition(&a, 0, n-1)

	fmt.Println(a)


}