class Solution {
    public boolean canBeEqual(int[] target, int[] arr) {
        
        Arrays.sort(target);
        Arrays.sort(arr);
        
        String st = Arrays.stream(target)
                        .mapToObj(String::valueOf)
                        .collect(Collectors.joining(""));
        
        String sA = Arrays.stream(arr)
                        .mapToObj(String::valueOf)
                        .collect(Collectors.joining(""));
        
        return sA.equals(st);
        
    }
}