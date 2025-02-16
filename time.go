package slimuuid

import (
	"time"
)
func TimePart() string {
	
	characters:= []rune("abABcdCD01efEFghGH23ijIJklKL45mnMNopOP67qrQRstST89uvUVwxWXyzYZ")

    // Define the reference time (Unix epoch for this example)
    referenceTime := time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC)

    // Get current time in UTC
    currentTime := time.Now().UTC()

    // Calculate the duration between reference time and current time
    duration := currentTime.Sub(referenceTime)

    // Calculate years, months, days, hours, minutes, and milliseconds
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
    milliseconds := duration.Milliseconds() % 1000 // Modulo by 60000 to get only milliseconds past the last minute
	
	timePart := string([]rune{
		characters[years],
		characters[months],
		characters[days],
		characters[hours],
		characters[minutes],
		characters[seconds],
		characters[int(milliseconds/62)],
		characters[int(milliseconds%62)],
	})

	return timePart
	
}

func TimePartFast() string {
	
	characters:= []rune("abABcdCD01efEFghGH23ijIJklKL45mnMNopOP67qrQRstST89uvUVwxWXyzYZ")

    // Define the reference time (Unix epoch for this example)
    referenceTime := time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC)

    // Get current time in UTC
    currentTime := time.Now().UTC()

    // Calculate the duration between reference time and current time
    duration := currentTime.Sub(referenceTime)

    // Calculate years, months, days, hours, minutes, and milliseconds
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
    milliseconds := duration.Milliseconds() % 1000 // Modulo by 60000 to get only milliseconds past the last minute
	nanoseconds := duration.Nanoseconds() % 1000
	timePart := string([]rune{
		characters[years],
		characters[months],
		characters[days],
		characters[hours],
		characters[minutes],
		characters[seconds],
		characters[int(milliseconds/62)],
		characters[int(milliseconds%62)],
        characters[int(nanoseconds/62)],
        characters[int(nanoseconds%62)],
	})

	return timePart
	
}