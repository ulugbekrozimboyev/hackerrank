import java.lang.Math;

class Solution {
    public boolean isPerfectSquare(int num) {
        for(int i = (int) Math.sqrt(num)-1; i*i <= num ; i++){
            if(i*i == num) return true;
            if(Math.pow(i, 2) > Integer.MAX_VALUE) break;
        }
        
        return false;
    }
}