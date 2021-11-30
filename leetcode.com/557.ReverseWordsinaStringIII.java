class Solution {
    public String reverseWords(String s) {
        StringBuilder sb = new StringBuilder();
        for(String part : s.split(" ")){
            sb.append(reverseString(part.toCharArray()));
            sb.append(" ");
        }
        
        return sb.toString().trim();
    }
    
    public String reverseString(char[] s) {
        int left = 0;
        int right = s.length - 1;
        while (left < right) {
            char temp = s[left];
            s[left++] = s[right];
            s[right--] = temp;
        }
        
        return new String(s);
    }
}