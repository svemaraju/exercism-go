package raindrops

import (
	"strconv"
)

func Convert(input int) string {
	// declare an empty string.
	raindrop := ""
	// append to the string when condition meets.
	if input%3 == 0 {
		raindrop += "Pling"
	}
	if input%5 == 0 {
		raindrop += "Plang"
	}
	if input%7 == 0 {
		raindrop += "Plong"
	}
	if len(raindrop) == 0 {
		raindrop = strconv.Itoa(input)
	}
	return raindrop
}
