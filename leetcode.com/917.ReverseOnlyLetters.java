class Solution {
    public String reverseOnlyLetters(String s) {
      StringBuilder sCopy = new StringBuilder();
      
      for(int i = s.length() -1; i >= 0; i--){
          if ((s.charAt(i) > 64 && s.charAt(i) < 91) || (s.charAt(i) > 96 && s.charAt(i) < 123)){
            sCopy.append(s.charAt(i));
          }
      }
      
      String sRevers = sCopy.toString();
      int j = 0;
      StringBuilder sResult = new StringBuilder();
      for(int i = 0; i < s.length(); i++){
        if (s.charAt(i) <= 64 || (s.charAt(i) >= 91 && s.charAt(i) <= 96 )){
            sResult.append(s.charAt(i));
        } else {
            sResult.append(sRevers.charAt(j++));
        }
      }

        
      return sResult.toString();
    }
}