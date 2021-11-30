import java.lang.Math;

class Solution {
    public double myPow(double x, int n) {
//         if(n == 0) return 1;
//         if(n == 1) return x;
        
//         if(n < 0) return 1 / (myPow(x, -1 * n));
        
//         return x * myPow(x, n-1);
        
        return Math.pow(x, n);
    }
}