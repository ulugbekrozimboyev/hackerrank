class Solution {
    public int[] shuffle(int[] nums, int n) {
        int[] newArray = new int[nums.length];
        
        for(int i = 0; i < n; i++) {
            newArray[i*2] = nums[i];
            newArray[i*2+1] = nums[n+i];
        }
        
        return newArray;
    }
}