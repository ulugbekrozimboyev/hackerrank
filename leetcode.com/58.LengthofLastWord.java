class Solution {
    public int lengthOfLastWord(String s) {
        s = s.trim();
        if (s.length() == 0 ) return 0;
        
        String[] arr = s.split(" ");
        
        return arr[arr.length-1].length();
    }
}