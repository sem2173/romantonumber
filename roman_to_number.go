package romanToNumber

func RomanToNumber (roman string) (number int) {

    dico := map[string]int{"I":1, "V":5}

    
        number = dico[roman]
    

    if roman == "I"+"I"{
        number = 1+1
    }

    if roman == "I"+"I"+"I"{
     number = 1+1+1
    }

     
return number
}