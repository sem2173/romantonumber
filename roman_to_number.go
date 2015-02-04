package romanToNumber

func RomanToNumber (roman string) (number int) {

if roman== "I"{
    number = 1
}

if roman== "I"+"I"{
    number = 1+1
}

if roman== "I"+"I"+"I"{
    number = 1+1+1
}

if roman== "V"{
    number = 5
}
 
return number
}