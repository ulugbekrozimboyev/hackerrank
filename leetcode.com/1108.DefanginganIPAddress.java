class Solution {
    public String defangIPaddr(String address) {
        
        String[] arr = address.split("\\.");
        
        return String.join("[.]", arr);
    }
}