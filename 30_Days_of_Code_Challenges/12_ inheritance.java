class Student extends Person{
	private int[] testScores;
    
    // Constructor
	Student(String firstName, String lastName, int identification, int[] testScores){
		super(firstName, lastName, identification);
        this.testScores = testScores;
	}
    
    public String calculate(){
        int z = 0;
        int len = this.testScores.length;
        String res;
        float avarage;
        for(int i = 0; i < len; i++){
            z += this.testScores[i];
        }
        avarage = z/len;
        if(90 <= avarage && avarage <= 100){ res = "O";}
        else if(80 <= avarage && avarage < 90 ){ res = "E";}
        else if( 70 <= avarage && avarage < 80) { res = "A"; }
        else if( 55 <= avarage && avarage < 70 ){ res = "P"; }
        else if( 40 <= avarage && avarage < 55){ res = "D";}
        else { res = "T"; }
        
        return res;
    }
}