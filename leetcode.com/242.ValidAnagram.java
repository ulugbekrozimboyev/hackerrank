import java.util.Arrays;

class Solution {
    public boolean isAnagram(String s, String t) {
        
        char[] cS = s.toCharArray();
        Arrays.sort(cS);
        String anagramOfS = new String(cS);
        
        char[] cT = t.toCharArray();
        Arrays.sort(cT);
        String anagramOfT = new String(cT);
        
        return anagramOfS.equals(anagramOfT);
    }
}