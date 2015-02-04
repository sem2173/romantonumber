package romanToNumber

func RomanToNumber (roman string) (number int) {

    dico := map[string]int{"I":1, "V":5}

    previousValue := 0

   for i := len(roman)- 1; i >=0; i--{
    currentLetter := string(roman[i])

    if i< (len(roman) -1) {
        previousValue = dico[string(roman[i+1])]
    }

    if previousValue > dico[currentLetter] {
        number -= dico[string(currentLetter)]
    } else {

        number += dico[string(currentLetter)]      
    }
   }     
    
return number
}