package slimuuid

import (
	"errors"
	"strconv"
)

func dateConverter(date string) (int, int, int, error) {
	year := date[:4]
	month := date[5:7]
	day := date[8:10]

    yearInt, err := strconv.Atoi(year)
    monthInt, err := strconv.Atoi(month)
    dayInt, err := strconv.Atoi(day)
    if err != nil {
        return 0, 0, 0, err
    }
	return yearInt, monthInt, dayInt, nil
}

// Generates a  slimuuid of 8(time chars)+ 12 (hash chars)  with non existent collison probability , so if you are using this with more than 100 ids per nano second this is the recommended function to use , it also utilizes the power of google uuid so if u want google's uuid robustness but with only 18 characters then use this function.
func Generate()  ( string , error ) {
    // generate a unique id using NewRandom function from /google/uuid package 
	uuid , err := ID()
    if err != nil {
        return "" , err 
    }

    // takes a time response clock synced upto millisecons encoded in 64 base character encoding
    timePart := NanoTime()

    // takes a hash of the uuid encoded in 64 base character encoding using murmur3 hash function
	hashedPart :=  SingleHashGenerator(uuid)

	return timePart + hashedPart , nil
}

// this function uses same algoritthm as Generate but uses a date of your choice from which nanoseconds will  be counted so you can send date in string format and it will generate a slimuuid based on that date , default  date in Generate function is 1 feb 2025 when slimuuid was first releaseddate date format should be like := "2025-02-01" / "2025/02/01"
func GenerateWithDate(date string) (string, error) {
    yearInt, monthInt, dayInt, err := dateConverter(date)
    if err != nil {
        return "" , err
    }

    uuid , err := ID()
    if err != nil {
        return "" , err 
    }
    timePart := NanoTimeWithDate(yearInt, monthInt, dayInt)
    hashedPart :=  SingleHashGenerator(uuid)

	return timePart + hashedPart , nil
}

// this function uses same algoritthm as Generate but uses a custom characters set of your choice , so you need to provide a characters set to the function that is a string of length 64 characters and try to make it as random as possible or use different characters like a-z , A-Z , 0-9 , etc. And Choose this string according to your requirement like while using this for genenerating. example : characters := "0123456789abcdefghijklmnopqrstuvwxyz_-ABCDEFGHIJKLMNOPQRSTUVWXYZ"
func GenerateWithCharacters(characters string) (string, error) {
    // check if characters is valid
    if len(characters) != 64 {
        return "" , errors.New("characters are not valid")
    }

     // generate a unique id using NewRandom function from /google/uuid package 
	uuid , err := ID()
    if err != nil {
        return "" , err 
    }

    // takes a time response clock synced upto millisecons encoded in 64 base character encoding with characters provided by you
    timePart := NanoTimeWithCharacters(characters)

    // takes a hash of the uuid encoded in 64 base character encoding using murmur3 hash function
	hashedPart :=  SingleHashGenerator(uuid)

    return timePart + hashedPart , nil
}

// same as GenerateWithCharacters function but with a date of your choice ,formats : characters := "0123456789abcdefghijklmnopqrstuvwxyz_-ABCDEFGHIJKLMNOPQRSTUVWXYZ" , date := "2025-02-01"
func GenerateWithCharactersAndDate(characters string, date string) (string, error) {
    if len(characters) != 64 {
        return "" , errors.New("characters are not valid")
    }
    yearInt, monthInt, dayInt, err := dateConverter(date)
    if err != nil {
        return "" , err
    }
    uuid , err := ID()
    if err != nil {
        return "" , err 
    }
    timePart := NanoTimeWithCharactersAndDate(characters, yearInt, monthInt, dayInt)
    hashedPart :=  SingleHashGenerator(uuid)
    return timePart + hashedPart , nil
}

// this is the best function to generate a slimuuid , it is the best because it is the most efficient and have least  characters that are 18 => 10 (time chars) + 8 (hash chars) . It takes a unique string and a counter that is used to  generate a unique id for each system , we recommend you use MAC address of the system as a unique string. which you can get using MacID function from unique.go file which is in the same package. 
func GenerateBest(unique string) (string, error) {
    timePart := NanoTime()
    hashedPart := SingleHashGenerator(unique + timePart) 
    return timePart + hashedPart, nil
}

// same as GenerateBest function but with a custom characters set of your choice , so you need to provide a characters set to the function that is a string of length 64 characters and try to make it as random as possible or use different characters like a-z , A-Z , 0-9 , etc. And Choose this string according to your requirement like while using this for genenerating. example : characters := "0123456789abcdefghijklmnopqrstuvwxyz_-ABCDEFGHIJKLMNOPQRSTUVWXYZ"
func GenerateBestWithCharacters(unique string, characters string) (string, error) {
    // check if characters is valid
    if len(characters) != 64 {
        return "" , errors.New("characters are not valid")
    }
    timePart := NanoTimeWithCharacters(characters)
	hashedPart :=  SingleHashGenerator(unique+timePart)
	return timePart + hashedPart , nil
}

// same as GenerateBest function but with a date of your choice , formats : date := "2025-02-01"
func GenerateBestWithDate(unique string, date string) (string, error) {
    yearInt, monthInt, dayInt, err := dateConverter(date)
    if err != nil {
        return "" , err
    }
    timePart := NanoTimeWithDate(yearInt, monthInt, dayInt)
	hashedPart :=  SingleHashGenerator(unique+timePart)
	return timePart + hashedPart , nil
}

func GenerateBestWithCharactersAndDate(unique string, characters string, date string) (string, error) {
    // check if characters is valid
    if len(characters) != 64 {
        return "" , errors.New("characters are not valid")
    }

    yearInt, monthInt, dayInt, err := dateConverter(date)
    if err != nil {
        return "" , err
    }

    timePart := NanoTimeWithCharactersAndDate(characters, yearInt, monthInt, dayInt)
    hashedPart :=  SingleHashGenerator(unique+timePart)
    return timePart + hashedPart , nil
    
}


