package main

import (
	"math"
	"strconv"
	"strings"
)

func ConvertBase(fromBase, toBase int, number string) string {
	var firstConversion int
	var finalConversion string

	switch fromBase {
	case 4:
		// This is for when the user is converting from decimal
		firstConversion, _ = strconv.Atoi(number)
	default:
		firstConversion = convertToDecimal(fromBase, number)
	}

	switch toBase {
	case 4:
		// This is for when the user is converting to decimal
		finalConversion = strconv.Itoa(firstConversion)
	default:
		finalConversion = convertFromDecimal(toBase, firstConversion)
	}

	return finalConversion
}

func convertToDecimal(fromBase int, number string) int {
	var base int
	var result int
	exponent := float64(len(number) - 1)

	switch fromBase {
	case 1:
		base = 2
	case 2:
		base = 8
	case 3:
		base = 16
	}

	for i := 0; i < len(number); i++ {
		if fromBase == 3 {
			number = strings.ToUpper(number)
			result += fromHexMap[string(number[i])] * int(math.Pow(float64(base), exponent))
			exponent--
			continue
		} else {
			current, _ := strconv.Atoi(string(number[i]))
			result += current * int(math.Pow(float64(base), exponent))
			exponent--
		}
	}
	return result
}

func convertFromDecimal(toBase int, decimal int) string {
	var base int
	var result string

	switch toBase {
	case 1:
		base = 2
	case 2:
		base = 8
	case 3:
		base = 16
	}

	for decimal > 0 {
		if toBase == 3 {
			result += toHexMap[decimal%base]
			decimal /= base
		} else {
			result += strconv.Itoa(decimal % base)
			decimal /= base
		}
	}

	// before returning the answer we must reverse the string, as it is currently
	// in reverse order: LSD -> MSD, we want MSD -> LSD
	return reverseString(result)
}

func isValidBase(fromBase int, number string) bool {

	if number == "" {
		return false
	}

	for i := 0; i < len(number); i++ {
		switch fromBase {
		case 1:
			if _, exists := BinMap[string(number[i])]; !exists {
				return false
			}
		case 2:
			if _, exists := OctMap[string(number[i])]; !exists {
				return false
			}
		case 3:
            number = strings.ToUpper(number)
			if _, exists := fromHexMap[string(number[i])]; !exists {
				return false
			}
		case 4:
			if _, err := strconv.Atoi(string(number[i])); err != nil {
				return false
			}
		}
	}

	return true
}

func reverseString(number string) string {
	byteArray := []byte(number)

	for i, j := 0, len(number)-1; i < j; i, j = i+1, j-1 {
		byteArray[i], byteArray[j] = byteArray[j], byteArray[i]
	}

	return string(byteArray)
}
