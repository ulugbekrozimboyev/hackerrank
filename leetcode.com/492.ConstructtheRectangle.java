class Solution {
    public int[] constructRectangle(int area) {
        int[] result = {area, 1}; 
        for (int i = (int) Math.sqrt(area); i <= area/2; i++) {
            if(area % i == 0 && i >= area/i) {
                result[0] = i;
                result[1] = area/i;
                return result; 
            }
        }
        
        return result;
    }
}