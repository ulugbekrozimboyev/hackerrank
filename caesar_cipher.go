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
    char* s = (char *)malloc(10240 * sizeof(char));
    scanf("%s",s);
    int k; 
    scanf("%d",&k);
    //s = "--/azZA";
    
    k = k%26;
    
    for(int i=0; i < n; i++ ) {
        
        
        if(s[i] > 64 && s[i] < 91) {
            int j = s[i] + k;
            (j > 90)? printf("%c",j-90+64) : printf("%c",j);
                
        } else if(s[i] > 96 && s[i] < 123) {
            int j = s[i] + k;
             
            (j > 122) ? printf("%c",j-122+96) : printf("%c",j);
        } else {            
            printf("%c",s[i]);            
        
        }
    }
    
    
    
    return 0;
}
