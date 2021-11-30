class Solution {
    public int romanToInt(String s) {
        int len = s.length() - 1;
        int result = 0;
        while(len >= 0) {
            int curValue = getValue(s.charAt(len));
            // System.out.println("curValue: " + curValue);
            int oldValue = result != 0 ? getValue(s.charAt(len+1)) : 0;
            // System.out.println("oldValue: "  + oldValue);
            if(result == 0 || oldValue <= curValue) {
                // System.out.println("plus + result: " + result + " curValue: " + curValue);
                result += curValue;
            } else {
                // System.out.println("minus - result: " + result + " curValue: " + curValue);
                result -= curValue;
                
            }
            len--;
        }
        return result;
    }
    
    public int getValue(char val){
        switch(val) {
            case 'I' : return 1;
            case 'V' : return 5;
            case 'X' : return 10;
            case 'L' : return 50;
            case 'C' : return 100;
            case 'D' : return 500;
            case 'M' : return 1000;
        }
        
        return 0;
    }
}