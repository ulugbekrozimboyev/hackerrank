class Solution {
    public boolean kLengthApart(int[] nums, int k) {
        
        int start = -1;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == 1) {
                // System.out.println(i);
                if (start == -1 ) { 
                    start = i;
                    // System.out.println(i);
                    continue;
                } else {
                    if (i - start -1 < k) return false;
                    start = i;
                    // System.out.println(i);
                }
            }
        }
        
        
        return true;
    }
}