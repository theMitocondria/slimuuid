package slimuuid

import (
	"time"
)

var characters = "abABcdCD01efEFghGH23-ijIJklKL45mnMNopOP67qrQRstS_T89uvUVwxWXyzYZ" 

//this function generates a time response clock synced upto nanoseconds encoded in 64 base character encoding
func NanoTime() string {
	referenceTime := time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC)
    currentTime := time.Now().UTC()
    duration := currentTime.Sub(referenceTime)
    years := currentTime.Year() - referenceTime.Year()
    months := int(currentTime.Month()) - int(referenceTime.Month())
    if months < 0 {
        years--
        months += 12
    }
    days := int(duration.Hours() / 24) % 30
    hours := int(duration.Hours()) % 24
    minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
    milliseconds := duration.Milliseconds() % 1000 
	nanoseconds := duration.Nanoseconds() % 1000
	return string(characters[years])+string(characters[months])+string(characters[days])+string(characters[hours])+string(characters[minutes])+string(characters[seconds])+string(characters[int(milliseconds/62)])+string(characters[int(milliseconds%62)])+string(characters[int(nanoseconds/62)])+string(characters[int(nanoseconds%62)]) 

}

//this function generates a time response clock synced upto nanoseconds encoded in 64 base character encoding with a custom characters set of your choice
func NanoTimeWithCharacters(chars string) string {
    referenceTime := time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC)
    currentTime := time.Now().UTC()
    duration := currentTime.Sub(referenceTime)
    years := currentTime.Year() - referenceTime.Year()
    months := int(currentTime.Month()) - int(referenceTime.Month())
    if months < 0 {
        years--
        months += 12
    }
    days := int(duration.Hours() / 24)
    hours := int(duration.Hours()) % 24
    minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
    milliseconds := duration.Milliseconds() % 1000 
	nanoseconds := duration.Nanoseconds() % 1000
	return string(chars[years])+string(chars[months])+string(chars[days])+string(chars[hours])+string(chars[minutes])+string(chars[seconds])+string(chars[int(milliseconds/62)])+string(chars[int(milliseconds%62)])+string(chars[int(nanoseconds/62)])+string(chars[int(nanoseconds%62)]) 

}

// this function generates a time response clock synced upto nanoseconds encoded in 64 base character encoding from a date given by you
func NanoTimeWithDate(year, month, day int) string {
    referenceTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
    currentTime := time.Now().UTC()
    duration := currentTime.Sub(referenceTime)
    years := currentTime.Year() - referenceTime.Year()
    months := int(currentTime.Month()) - int(referenceTime.Month())
    if months < 0 {
        years--
        months += 12
    }
    days := int(duration.Hours() / 24)
    hours := int(duration.Hours()) % 24
    minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
    milliseconds := duration.Milliseconds() % 1000 
	nanoseconds := duration.Nanoseconds() % 1000
	return string(characters[years])+string(characters[months])+string(characters[days])+string(characters[hours])+string(characters[minutes])+string(characters[seconds])+string(characters[int(milliseconds/62)])+string(characters[int(milliseconds%62)])+string(characters[int(nanoseconds/62)])+string(characters[int(nanoseconds%62)]) 

}

// this function generates a time response clock synced upto nanoseconds encoded in 64 base character encoding with a custom characters set of your choice and a date given by you
func NanoTimeWithCharactersAndDate(chars string, year, month, day int) string {
    referenceTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
    currentTime := time.Now().UTC()
    duration := currentTime.Sub(referenceTime)
    years := currentTime.Year() - referenceTime.Year()
    months := int(currentTime.Month()) - int(referenceTime.Month())
    if months < 0 {
        years--
        months += 12
    }
    days := int(duration.Hours() / 24)
    hours := int(duration.Hours()) % 24
    minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
    milliseconds := duration.Milliseconds() % 1000 
	nanoseconds := duration.Nanoseconds() % 1000
	return string(chars[years])+string(chars[months])+string(chars[days])+string(chars[hours])+string(chars[minutes])+string(chars[seconds])+string(chars[int(milliseconds/62)])+string(chars[int(milliseconds%62)])+string(chars[int(nanoseconds/62)])+string(chars[int(nanoseconds%62)]) 
}
