package romanToNumber

func RomanToNumber (roman string) (number int) {

    dico := map[string]int{"I":1, "V":5}

    
        number = dico[roman]
    

    if roman == "I"+"I"{
        number = dico["I"] + dico["I"]
    }

    if roman == "I"+"I"+"I"{
     number = dico["I"] + dico["I"] + dico["I"]
    }

     
return number
}