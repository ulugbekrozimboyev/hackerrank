class Solution {
    public int maxProduct(int[] nums) {
        int result = 0;
        for(int i = 0; i < nums.length-1; i++){
            for(int j = i+1; j < nums.length; j++){
                if (result < (nums[i] - 1) * (nums[j] - 1)) {
                    result = (nums[i] - 1) * (nums[j] - 1);
                }
            }
        }
        
        return result;
    }
}