class Solution {
    public String reverseStr(String s, int k) {
        
        char[] arr = s.toCharArray();
       
        
        for (int i = 0; i < arr.length; i+= 2*k) {
            int start = i, end = Math.min(i + k - 1, arr.length - 1);
            while (start < end) {
                char temp = arr[start];
                arr[start++] = arr[end];
                arr[end--] = temp;                
            }
        }
        
        // System.out.print(new String(arr));
        
        return new String(arr);
    }
}