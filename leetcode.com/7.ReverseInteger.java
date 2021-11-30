class Solution {
    public int reverse(int x) {
        
        // System.out.println(" x = " + x);
        int result = 0;
        while(x != 0){
            // System.out.println(" x = " + x);
            int additional = x % 10;
            x = (x-additional)/10;

            // System.out.println(" result, additional = " + result + "-" + additional);
            
            if(result > Integer.MAX_VALUE/10 || result < Integer.MIN_VALUE/10 ) return 0;
            
            result = result * 10 + additional;
        }
        // System.out.println(" result = " + result );
        return result;
    }
}