package romanToNumber

func RomanToNumber (roman string) (number int) {

    dico := map[string]int{"I":1, "V":5}

    if roman== "I"{
        number = dico["I"]
    }

    if roman== "I"+"I"{
        number = 1+1
    }

    if roman== "I"+"I"+"I"{
     number = 1+1+1
    }

    if roman== "V"{
        number = dico["V"]
    }
 
return number
}