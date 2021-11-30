class Solution {
    
    public boolean contansArr(String tmpPrefix, String[] strs){
        // System.out.println("contains arr");
        for(int i = 1; i < strs.length; i++){
            // System.out.println(strs[i]);    
            if (!strs[i].startsWith(tmpPrefix)) {
                return false;
            }

        }
        
        return true;
    }
    
    public String longestCommonPrefix(String[] strs) {
        
        StringBuilder tmpPrefix = new StringBuilder(strs[0]);
        
        while(tmpPrefix.length() > 0) {
            // System.out.println(tmpPrefix);
            if( contansArr(tmpPrefix.toString(), strs) ) break;
            
            tmpPrefix.deleteCharAt(tmpPrefix.length()-1); 
        }
        
        return tmpPrefix.toString();
    }
}