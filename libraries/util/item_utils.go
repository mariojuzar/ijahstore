package util

import (
	"fmt"
	"strconv"
	"strings"
)

func ShortenSizeStr(oldSize string) string {
	var newSize = ""

	switch oldSize {
	case "XXXL":
		newSize = "X3"
	case "XXL":
		newSize = "XX"
	case "XL":
		newSize = "XL"
	default:
		newSize = oldSize + oldSize
	}

	return newSize
}

func ShortenColourStr(colour string) string {
	newColour := strings.ToUpper(colour[0:3])

	if strings.Contains(colour, " ") {
		slice := strings.Split(colour, " ")
		newColour = strings.ToUpper(slice[0][0:1] + slice[1][0:2])
	}

	return newColour
}

func GenerateSKU(ID uint, size string, colour string) string {
	return "SSI-D" + PrettifySKUIDToString(ID) + "-" + ShortenSizeStr(size) + "-" + ShortenColourStr(colour)
}

func StrToUint(word string) uint {

	num, _ := strconv.Atoi(word)
	return uint(num)
}

func PrettifySKUIDToString(id uint) string  {
	return fmt.Sprintf("%08d", id)
}