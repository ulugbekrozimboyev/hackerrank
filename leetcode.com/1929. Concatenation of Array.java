class Solution {
    public int[] getConcatenation(int[] nums) {
        int length = nums.length;
        int[] result = new int[2 * length];
        System.arraycopy(nums, 0, result, 0, length);
        System.arraycopy(nums, 0, result, length, length);
        
        return result;
    }
}