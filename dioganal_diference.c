#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

int main(){
    int n; 
    scanf("%d",&n);
    int a[n][n];
    long int s1 =0, s2=0;
    long int result;
    for(int a_i = 0; a_i < n; a_i++){
       for(int a_j = 0; a_j < n; a_j++){
          
          scanf("%d",&a[a_i][a_j]);
       }
       s1 += a[a_i][a_i];
       s2 += a[a_i][n-a_i-1];
    }
    
    printf("%ld", abs(s1-s2));
    
    return 0;
}
