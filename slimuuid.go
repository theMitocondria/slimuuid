package slimuuid

import (
	"errors"
	"strconv"
    "time"
    "sync/atomic"
)

var counter uint64 = 0

/* Generates a  slimuuid of 8(time chars)+ 12 (hash chars)  with non existent collison probability , so if you are using
    this with more than 100 ids per nano second this is the recommended function to use . 
*/
func Generate()  ( string , error ) {
    // generate a unique id using NewRandom function from /google/uuid package 
	uuid , err := ID()
    if err != nil {
        return "" , err 
    }

    // takes a time response clock synced upto millisecons encoded in 64 base character encoding
    timePart := MilliTime()

    // takes a hash of the uuid encoded in 64 base character encoding using murmur3 hash function
	hashedPart :=  DoubleHashGenerator(uuid)

    // combines the time part and the hashed part to form a slimuuid
	slimId := timePart + hashedPart 

	return slimId , nil
}

// same as Generate function but with a seed , so you need to provide a seed to the function that is a uint32
func GenerateWithSeed(seed uint32) (string, error) {
    // check if seed is valid
    if seed < 0 {
        return "" , errors.New("seed is not valid")
    }

    //check if seed is 0
    if seed == 0 {
        seed = uint32(time.Now().UnixNano())
    }

    // generate a unique id using NewRandom function from /google/uuid package 
	uuid , err := ID()
    if err != nil {
        return "" , err 
    }

    

    // takes a time response clock synced upto millisecons encoded in 64 base character encoding
    timePart := MilliTime()

    // takes a hash of the uuid encoded in 64 base character encoding using murmur3 hash function
	hashedPart :=  DoubleHashGeneratorWithSeed(uuid , seed)

    // combines the time part and the hashed part to form a slimuuid
	slimId := timePart + hashedPart 

	return slimId , nil
}

/* 
    this function uses same algoritthm as Generate but uses a date of your choice from which milliseconds will
    be counted so you can send date in string format and it will generate a slimuuid based on that date , default 
    date in Generate function is 1 feb 2025 when slimuuid was first released.
    always call the function like => 
    date := "2025-02-01" / "2025/02/01"
    slimId := GenerateWithDate(date)
    fmt.Println(slimId)
    Always pass the date in string format and not a time.Time format 
*/
func GenerateWithDate(date string) (string, error) {
    // take out year , month , day from string
    year := date[:4]
    month := date[5:7]
    day := date[8:10]

    // convert year , month , day to int
    yearInt, err := strconv.Atoi(year)
    monthInt, err := strconv.Atoi(month)
    dayInt, err := strconv.Atoi(day)
    if err != nil {
        return "" , err
    }
    
    uuid , err := ID()
    if err != nil {
        return "" , err 
    }

    // takes a time response clock synced upto milliseconds encoded in 64 base character encoding
    timePart := MilliTimeWithDate(yearInt, monthInt, dayInt)

    // takes a hash of the uuid encoded in 64 base character encoding using murmur3 hash function
	hashedPart :=  DoubleHashGenerator(uuid)

    // combines the time part and the hashed part to form a slimuuid
	slimId := timePart + hashedPart 

	return slimId , nil
}

/*
    same as GenerateWithDate function but with a seed , so you need to provide a seed to the function that is a uint32
    always call the function like => 
    date := "2025-02-01"
    seed := uint32(time.Now().UnixNano()) or any integer you want to use as a seed greater than 0
    slimId := GenerateWithDateAndSeed(date, seed)
    fmt.Println(slimId)
    Always pass the date in string format and not a time.Time format  
*/
func GenerateWithDateAndSeed(date string, seed uint32) (string, error) {
    // take out year , month , day from string
    year := date[:4]
    month := date[5:7]
    day := date[8:10]

    // convert year , month , day to int
    yearInt, err := strconv.Atoi(year)
    monthInt, err := strconv.Atoi(month)
    dayInt, err := strconv.Atoi(day)
    if err != nil {
        return "" , err
    }

    // check if seed is valid
    if seed < 0 {
        return "" , errors.New("seed is not valid")
    }

    //check if seed is 0
    if seed == 0 {
        seed = uint32(time.Now().UnixNano())
    }

    uuid , err := ID()
    if err != nil {
        return "" , err 
    }

    // takes a time response clock synced upto milliseconds encoded in 64 base character encoding
    timePart := MilliTimeWithDate(yearInt, monthInt, dayInt)

    // takes a hash of the uuid encoded in 64 base character encoding using murmur3 hash function
	hashedPart :=  DoubleHashGeneratorWithSeed(uuid , seed)

    // combines the time part and the hashed part to form a slimuuid
	slimId := timePart + hashedPart 

    return slimId , nil
}

/*
    this function uses same algoritthm as Generate but uses a custom characters set of your choice , so you need to provide a characters set to the function that is a string
    always call the function like => 
    characters := "0123456789abcdefghijklmnopqrstuvwxyz_-ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    slimId := GenerateWithCharacters(characters)
    fmt.Println(slimId)
    Always pass the string of length 64 characters and try to make it as random as possible or use different characters like a-z , A-Z , 0-9 , etc. And Choose this string according to your requirement like while using this for genenerating file name 
    you can not use some characters in file name like / , \ , : , * , ? , " , < , > , | , etc.
*/
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
    timePart := MilliTimeWithCharacters(characters)

    // takes a hash of the uuid encoded in 64 base character encoding using murmur3 hash function
	hashedPart :=  DoubleHashGenerator(uuid)

    // combines the time part and the hashed part to form a slimuuid
	slimId := timePart + hashedPart 

	return slimId , nil
}

/*
    same as GenerateWithCharacters function but with a seed , so you need to provide a seed to the function that is a uint32
    always call the function like => 
    characters := "0123456789abcdefghijklmnopqrstuvwxyz_-ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    seed := uint32(time.Now().UnixNano()) or any integer you want to use as a seed greater than 0
    slimId := GenerateWithCharactersAndSeed(characters, seed)
    fmt.Println(slimId)
    Always pass the string of length 64 characters and try to make it as random as possible or use different characters like a-z , A-Z , 0-9 , etc. And Choose this string according to your requirement like while using this for genenerating file name 
    you can not use some characters in file name like / , \ , : , * , ? , " , < , > , | , etc.
*/
func GenerateWithCharactersAndSeed(characters string, seed uint32) (string, error) {
    // check if characters is valid
    if len(characters) != 64 {
        return "" , errors.New("characters are not valid")
    }

    // check if seed is valid
    if seed < 0 {
        return "" , errors.New("seed is not valid")
    }

    //check if seed is 0
    if seed == 0 {
        seed = uint32(time.Now().UnixNano())
    }

    // generate a unique id using NewRandom function from /google/uuid package 
	uuid , err := ID()
    if err != nil {
        return "" , err 
    }

    

    // takes a time response clock synced upto millisecons encoded in 64 base character encoding with characters provided by you
    timePart := MilliTimeWithCharacters(characters)

    // takes a hash of the uuid encoded in 64 base character encoding using murmur3 hash function
	hashedPart :=  DoubleHashGeneratorWithSeed(uuid , seed)

    // combines the time part and the hashed part to form a slimuuid
	slimId := timePart + hashedPart 

	return slimId , nil
    
}

/*
    same as GenerateWithCharacters function but with a date of your choice , so you need to provide a date to the function that is a string
    always call the function like => 
    characters := "0123456789abcdefghijklmnopqrstuvwxyz_-ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    date := "2025-02-01"
    slimId := GenerateWithCharactersAndDate(characters, date)
    fmt.Println(slimId)
*/
func GenerateWithCharactersAndDate(characters string, date string) (string, error) {
    // check if characters is valid
    if len(characters) != 64 {
        return "" , errors.New("characters are not valid")
    }

    // take out year , month , day from string
    year := date[:4]
    month := date[5:7]
    day := date[8:10]

    // convert year , month , day to int
    yearInt, err := strconv.Atoi(year)
    monthInt, err := strconv.Atoi(month)
    dayInt, err := strconv.Atoi(day)
    if err != nil {
        return "" , err
    }
    
    uuid , err := ID()
    if err != nil {
        return "" , err 
    }

    // takes a time response clock synced upto milliseconds encoded in 64 base character encoding with characters provided by you
    timePart := MilliTimeWithCharactersAndDate(characters, yearInt, monthInt, dayInt)

    // takes a hash of the uuid encoded in 64 base character encoding using murmur3 hash function
    hashedPart :=  DoubleHashGenerator(uuid)

    // combines the time part and the hashed part to form a slimuuid
    slimId := timePart + hashedPart 

    return slimId , nil
}

/*
    same as GenerateWithCharactersAndDate function but with a seed , so you need to provide a seed to the function that is a uint32
    always call the function like => 
    characters := "0123456789abcdefghijklmnopqrstuvwxyz_-ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    date := "2025-02-01"
    seed := uint32(time.Now().UnixNano()) or any integer you want to use as a seed greater than 0
    slimId := GenerateWithCharactersAndDateAndSeed(characters, date, seed)
    fmt.Println(slimId)
    Keep the format of date right , characters should be of length 64 and seed should be greater than 0 
*/
func GenerateWithCharactersAndDateAndSeed(characters string, date string, seed uint32) (string, error) {
    // check if characters is valid
    if len(characters) != 64 {
        return "" , errors.New("characters are not valid")
    }

    // take out year , month , day from string
    year := date[:4]
    month := date[5:7]
    day := date[8:10]

    // convert year , month , day to int
    yearInt, err := strconv.Atoi(year)
    monthInt, err := strconv.Atoi(month)
    dayInt, err := strconv.Atoi(day)
    if err != nil {
        return "" , err
    }

    // check if seed is valid
    if seed < 0 {
        return "" , errors.New("seed is not valid")
    }

    //check if seed is 0
    if seed == 0 {
        seed = uint32(time.Now().UnixNano())
    }

    uuid , err := ID()
    if err != nil {
        return "" , err 
    }

    // takes a time response clock synced upto milliseconds encoded in 64 base character encoding
    timePart := MilliTimeWithCharactersAndDate(characters, yearInt, monthInt, dayInt)

    // takes a hash of the uuid encoded in 64 base character encoding using murmur3 hash function
	hashedPart :=  DoubleHashGeneratorWithSeed(uuid , seed)

    // combines the time part and the hashed part to form a slimuuid
	slimId := timePart + hashedPart 

    return slimId , nil
    
}

/* this is the best function to generate a slimuuid , it is the best because it is the most efficient and have least 
  characters that are 18 => 10 (time chars) + 8 (hash chars) . It takes a unique string and a counter that is used to 
  generate a unique id for each system , we recommend you use MAC address of the system as a unique string. which you
  can get using MacID function from unique.go file which is in the same package. 
  We could have integrated generating it ourself but it would make a significant downgrade in performance , so u just 
  generate in once then you are good to go. 
  always call the function like => 
  MacID , err := MacID()
  if err != nil {
    return "" , err
  }
  and store your MAC ID in a variable and pass it to the function like => 
  unique := "your_mac_id"
  slimId := GenerateBest(unique)
  and then you can use it like => 
  fmt.Println(slimId)
*/
func GenerateBest(unique string) string {
    /*
     takes a hash of the unique string provided by you + counter of thread in a Nanosecond 
     so almost non existent collision probability + time part encoded in 64 base character encoding 
     using murmur3 hash function
    */
    // Atomically increment and get counter
    curr := atomic.AddUint64(&counter, 1)
    
    // Generate hasher string
    hasher := ""
    for curr > 0 {
        hasher = string(characters[curr%64]) + hasher
        curr = curr / 64
    }

    timePart := NanoTime()
    hashedPart := SingleHashGenerator(unique + hasher + timePart)
    return timePart + hashedPart
}


func GenerateBestWithCharacters(unique string, characters string) (string, error) {
    // check if characters is valid
    if len(characters) != 64 {
        return "" , errors.New("characters are not valid")
    }
    
    // takes a time response clock synced upto nanoseconds encoded in 64 base character encoding
    timePart := NanoTimeWithCharacters(characters)

    /*
     takes a hash of the unique string provided by you + counter of thread in a Nanosecond 
     so almost non existent collision probability + time part encoded in 64 base character encoding 
     using murmur3 hash function
    */
	hashedPart :=  SingleHashGenerator(unique+string(characters[counter])+timePart)

    // if counter is greater than 64 then reset it to 0
    if counter >= 64 {
        counter = 0
    }

    // increment the counter
	return timePart + hashedPart , nil
    
}
/*
  same as GenerateBest function but with a seed , so you need to provide a seed to the function that is a uint32
  always call the function like => 
  MacID , err := MacID()
  if err != nil {
    return "" , err
  }
  and store your MAC ID in a variable and pass it to the function like => 
  unique := "your_mac_id"
  seed := uint32(time.Now().UnixNano()) or any integer you want to use as a seed greater than 0
  slimId, err := GenerateBestWithSeed(unique, seed)
  if err != nil {
    return "" , err
  }
  fmt.Println(slimId)
  Always pass the seed as a uint32 and not a int32 and do remember seed creates a better hash but is 10-15ns slower
*/
func GenerateBestWithSeed(unique string, seed uint32) (string, error) {
    // check if seed is valid
    if seed < 0 {
        return "" , errors.New("seed is not valid")
    }

    //check if seed is 0
    if seed == 0 {
        seed = uint32(time.Now().UnixNano())
    }
    
    // takes a time response clock synced upto nanoseconds encoded in 64 base character encoding
    timePart := NanoTime()

    /*
     takes a hash of the unique string provided by you + counter of thread in a Nanosecond 
     so almost non existent collision probability + time part encoded in 64 base character encoding 
     using murmur3 hash function
    */
	hashedPart :=  SingleHashGeneratorWithSeed(unique+string(characters[counter])+timePart, seed)

    // if counter is greater than 64 then reset it to 0
    if counter >= 64 {
        counter = 0
    }

    // increment the counter
	return timePart + hashedPart , nil
}

/*
  same as GenerateBestWithCharacters function but with a seed , so you need to provide a seed to the function that is a uint32
  always call the function like => 
  characters := "0123456789abcdefghijklmnopqrstuvwxyz_-ABCDEFGHIJKLMNOPQRSTUVWXYZ"
  unique := "your_mac_id"
  seed := uint32(time.Now().UnixNano()) or any integer you want to use as a seed greater than 0
  slimId, err := GenerateBestWithCharactersAndSeed(unique, characters, seed)
  if err != nil {
    return "" , err
  }
  fmt.Println(slimId)
*/
func GenerateBestWithCharactersAndSeed(unique string, characters string, seed uint32) (string, error) {
    // check if characters is valid
    if len(characters) != 64 {
        return "" , errors.New("characters are not valid")
    }

     // check if seed is valid
     if seed < 0 {
        return "" , errors.New("seed is not valid")
    }

    //check if seed is 0
    if seed == 0 {
        seed = uint32(time.Now().UnixNano())
    }
    
    // takes a time response clock synced upto nanoseconds encoded in 64 base character encoding
    timePart := NanoTimeWithCharacters(characters)

    /*
     takes a hash of the unique string provided by you + counter of thread in a Nanosecond 
     so almost non existent collision probability + time part encoded in 64 base character encoding 
     using murmur3 hash function
    */
	hashedPart :=  SingleHashGeneratorWithSeed(unique+string(characters[counter])+timePart, seed)

    // if counter is greater than 64 then reset it to 0
    if counter >= 64 {
        counter = 0
    }

    // increment the counter
	return timePart + hashedPart , nil
    
}

/*
  same as GenerateBest function but with a date of your choice , so you need to provide a date to the function that is a string
  always call the function like => 
  date := "2025-02-01"
  unique := "your_mac_id"
  slimId, err := GenerateBestWithDate(unique, date)
  if err != nil {
    return "" , err
  }
  fmt.Println(slimId)
  Always pass the date in string format and not a time.Time format  
*/
func GenerateBestWithDate(unique string, date string) (string, error) {
    // take out year , month , day from string
    year := date[:4]
    month := date[5:7]
    day := date[8:10]
    
    // convert year , month , day to int
    yearInt, err := strconv.Atoi(year)
    monthInt, err := strconv.Atoi(month)
    dayInt, err := strconv.Atoi(day)
    if err != nil {
        return "" , err
    }

    // takes a time response clock synced upto nanoseconds encoded in 64 base character encoding from date given by you
    timePart := NanoTimeWithDate(yearInt, monthInt, dayInt)

    /*
     takes a hash of the unique string provided by you + counter of thread in a Nanosecond 
     so almost non existent collision probability + time part encoded in 64 base character encoding 
     using murmur3 hash function
    */
	hashedPart :=  SingleHashGenerator(unique+string(characters[counter])+timePart)

    // if counter is greater than 64 then reset it to 0
    if counter >= 64 {
        counter = 0
    }

    // increment the counter
	return timePart + hashedPart , nil
}

func GenerateBestWithCharactersAndDate(unique string, characters string, date string) (string, error) {
    // check if characters is valid
    if len(characters) != 64 {
        return "" , errors.New("characters are not valid")
    }

    // take out year , month , day from string
    year := date[:4]
    month := date[5:7]
    day := date[8:10]
    
    // convert year , month , day to int
    yearInt, err := strconv.Atoi(year)
    monthInt, err := strconv.Atoi(month)
    dayInt, err := strconv.Atoi(day)
    if err != nil {
        return "" , err
    }

    // takes a time response clock synced upto nanoseconds encoded in 64 base character encoding from date given by you
    timePart := NanoTimeWithCharactersAndDate(characters, yearInt, monthInt, dayInt)

    /*
    takes a hash of the unique string provided by you + counter of thread in a Nanosecond 
    so almost non existent collision probability + time part encoded in 64 base character encoding 
    using murmur3 hash function
    */
    hashedPart :=  SingleHashGenerator(unique+string(characters[counter])+timePart)

    // if counter is greater than 64 then reset it to 0
    if counter >= 64 {
        counter = 0
    }

    // increment the counter
    return timePart + hashedPart , nil
    
}

/*
  same as GenerateBestWithDate function but with a seed , so you need to provide a seed to the function that is a uint32
  always call the function like => 
  date := "2025-02-01"
  unique := "your_mac_id"
  seed := uint32(time.Now().UnixNano()) or any integer you want to use as a seed greater than 0
  slimId, err := GenerateBestWithDateAndSeed(unique, date, seed)
  if err != nil {
    return "" , err
  }
  fmt.Println(slimId)
  Always pass the date in string format and not a time.Time format  
*/
func GenerateBestWithDateAndSeed(unique string, date string, seed uint32) (string, error) {    
    // take out year , month , day from string
    year := date[:4]
    month := date[5:7]
    day := date[8:10]
    
    // convert year , month , day to int
    yearInt, err := strconv.Atoi(year)
    monthInt, err := strconv.Atoi(month)
    dayInt, err := strconv.Atoi(day)
    if err != nil {
        return "" , err
    }   

    // check if seed is valid
    if seed < 0 {
        return "" , errors.New("seed is not valid")
    }
    
    // check if seed is 0
    if seed == 0 {
        seed = uint32(time.Now().UnixNano())
    }

    // takes a time response clock synced upto nanoseconds encoded in 64 base character encoding from date given by you
    timePart := NanoTimeWithDate(yearInt, monthInt, dayInt)

    /*
     takes a hash of the unique string provided by you + counter of thread in a Nanosecond 
     so almost non existent collision probability + time part encoded in 64 base character encoding 
     using murmur3 hash function
    */
    hashedPart :=  SingleHashGeneratorWithSeed(unique+string(characters[counter])+timePart, seed)

    // if counter is greater than 64 then reset it to 0
    if counter >= 64 {
        counter = 0
    }

    // increment the counter
    return timePart + hashedPart , nil
}

func GenerateBestWithCharactersAndDateAndSeed(unique string, characters string, date string, seed uint32) (string, error) {
    // check if characters is valid
    if len(characters) != 64 {
        return "" , errors.New("characters are not valid")
    }

    // take out year , month , day from string
    year := date[:4]
    month := date[5:7]
    day := date[8:10]
    
    // convert year , month , day to int
    yearInt, err := strconv.Atoi(year)
    monthInt, err := strconv.Atoi(month)
    dayInt, err := strconv.Atoi(day)
    if err != nil {
        return "" , err
    }   

    // check if seed is valid
    if seed < 0 {
        return "" , errors.New("seed is not valid")
    }
    
    // check if seed is 0
    if seed == 0 {
        seed = uint32(time.Now().UnixNano())
    }

    // takes a time response clock synced upto nanoseconds encoded in 64 base character encoding from date given by you
    timePart := NanoTimeWithCharactersAndDate(characters, yearInt, monthInt, dayInt)

    /*
     takes a hash of the unique string provided by you + counter of thread in a Nanosecond 
     so almost non existent collision probability + time part encoded in 64 base character encoding 
     using murmur3 hash function
    */
    hashedPart :=  SingleHashGeneratorWithSeed(unique+string(characters[counter])+timePart, seed)

    // if counter is greater than 64 then reset it to 0
    if counter >= 64 {
        counter = 0
    }

    // increment the counter
    return timePart + hashedPart , nil 
}


