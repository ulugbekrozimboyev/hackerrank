class Solution {
    public String restoreString(String s, int[] indices) {
        
        char[] arrS = s.toCharArray();
        
        StringBuilder sb = new StringBuilder();
        for(int i = 0; i < indices.length; i++){
            sb.append(arrS[getIndexOf(indices, i)]);
        }
        
        return sb.toString();
    }
    
    private int getIndexOf(int[] nums, int val){
        for(int i = 0; i < nums.length; i++) {
            if(nums[i] == val) return i;
        }
        
        return -1;
    }
}