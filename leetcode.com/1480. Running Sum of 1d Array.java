class Solution {
    public int[] runningSum(int[] nums) {
        int[] newArr = new int[nums.length];
        int sum = 0;
        for(int i = 0; i < nums.length; i++){
            sum += nums[i];
            newArr[i] = sum;
            // System.out.println(sum);
        }
        
        return newArr;
    }
}