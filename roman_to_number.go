package romanToNumber

func RomanToNumber (roman string) (number int) {

    dico := map[string]int{"I":1, "V":5}

   for _, romanletter :=range roman {
        number += dico[string(romanletter)]
    }     
    
return number
}