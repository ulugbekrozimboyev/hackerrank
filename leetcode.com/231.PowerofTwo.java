class Solution {
    public boolean isPowerOfTwo(int n) {
        if (n == 0) return false;
        if (n == 1) return true;
        
        return (n % 2 == 1) ? false : isPowerOfTwo(n/2);
    }
}