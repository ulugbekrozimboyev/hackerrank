class Solution {
    public List<String> fizzBuzz(int n) {
        
        List<String> list = new ArrayList<>();
        IntStream.range(1, n+1).forEach(item -> list.add(getType(item)));
        
        return list;
    }
    
    private String getType(int n) {
        if (n % 15 == 0 ) return "FizzBuzz";
        if (n % 3 == 0) return "Fizz";
        if (n % 5 == 0) return "Buzz";
        
        return String.valueOf(n);
    }
}