class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        int m = matrix.length;
        int n = matrix[0].length;
        List<Integer> list = new ArrayList();
        
        int left = 0, top = 0;
        int right = n-1, bottom = m-1;
        int direction = 0;
        
        while(left <= right && top <= bottom) {
            if(direction == 0) {
                for(int i = left; i <= right; i++){
                    list.add(matrix[top][i]);
                }
                top++;
            }
            
            if(direction == 1) {
                for(int i = top; i <= bottom; i++){
                    list.add(matrix[i][right]);
                }
                right--;
            }
            
            if(direction == 2) {
                for(int i = right; i >= left; i--){
                    list.add(matrix[bottom][i]);
                }
                bottom--;
            }
            
            if(direction == 3) {
                for(int i = bottom; i >= top; i--){
                    list.add(matrix[i][left]);
                }
                left++;
            }
        
            direction = (direction + 1) % 4;
        }
        
        
        return list;
    }
}