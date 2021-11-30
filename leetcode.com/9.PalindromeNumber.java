class Solution {
    public boolean isPalindrome(int x) {
        
        if (x < 0 || x != reverse(x)) {
            return false;
        }
        
        return true;
    }
    
    public int reverse(int x) {
        
        int result = 0;
        while(x != 0){
            int additional = x % 10;
            x = (x-additional)/10;
            
            if(result > Integer.MAX_VALUE/10 || result < Integer.MIN_VALUE/10 ) return 0;
            
            result = result * 10 + additional;
        }
        return result;
    }
}