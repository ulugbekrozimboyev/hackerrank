class Solution {
    public String truncateSentence(String s, int k) {
        String[] arr = s.split(" ");
        return String.join(" ", Arrays.copyOfRange(arr, 0, k));
    }
}