package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func IsValidCoordinates(coordinates string) bool {
	var longlat = strings.Split(coordinates, ",")

	// comma rule
	if len(longlat) == 2 && len(longlat[0]) > 0 && len(longlat[1]) > 0 {
		var longitude = strings.TrimSpace(longlat[1])
		var latitude = strings.TrimSpace(longlat[0])

		// format rule
		var pattern = "^-?\\d+(?:\\.?\\d+)?$"
		var longitudeMatched, _ = regexp.MatchString(pattern, longitude)
		var latitudeMatched, _ = regexp.MatchString(pattern, latitude)

		if longitudeMatched && latitudeMatched {
			// range rule
			var longitudeF, _ = strconv.ParseFloat(longitude, 10)
			var latitudeF, _ = strconv.ParseFloat(latitude, 10)

			if latitudeF >= -90.0 && latitudeF <= 90.0 && longitudeF >= -180.0 && longitudeF <= 180.0 {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(IsValidCoordinates("-23, 25"))
	fmt.Println(IsValidCoordinates("4, -3"))
	fmt.Println(IsValidCoordinates("-24.53525235, 23.45235"))
	fmt.Println(IsValidCoordinates("04, -23.234235"))
	fmt.Println(IsValidCoordinates("43.91343345, 143"))
}
