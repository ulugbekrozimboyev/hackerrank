import java.text.DateFormat;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.Calendar;
import java.text.ParseException;

class Solution {
    
    private final DateFormat formatter = new SimpleDateFormat("yyyy-MM-dd");
    
    public int dayOfYear(String date) {
    
        
        int days = 0;
        
        // System.out.println(date);
        
        try {
            Date endDate = formatter.parse(date);
            
            // System.out.println(endDate);
            
            Calendar cal = Calendar.getInstance();
            cal.setTime(endDate);
            cal.set(Calendar.DAY_OF_YEAR, 1);
            Date startDate = cal.getTime();
            
            // System.out.println(startDate);
            
            long diff = endDate.getTime() - startDate.getTime();
            days =(int) (diff / 1000 / 60 / 60 / 24) + 1;
            // System.out.println ("Days: " + days);
            
            
        } catch (ParseException excpt) {
            excpt.printStackTrace();
        }
        
        return days;
    }
}