	
public Difference(int[] elements){
    this.elements = elements;
}

public void computeDifference(){
    maximumDifference = 0;
    for(int i=0; i< elements.length; i++){
        for(int j = i; j < elements.length; j++){
            if(Math.abs(elements[i]-elements[j]) > maximumDifference) {
                maximumDifference = Math.abs(elements[i]-elements[j]);
            }
        }
    }
}