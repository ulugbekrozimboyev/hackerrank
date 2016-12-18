//Write MyBook Class
class MyBook extends Book {
    private int price;
    
    MyBook(String t,String a,int price){
        super(t,a);
        this.price = price;
    }
    
    public void display() {
            System.out.println("Title: " + this.title
            + "\nAuthor: " + this.author
            + "\nPrice: " + this.price);
    }
    
}