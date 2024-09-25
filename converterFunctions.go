package main

import (
	"math"
	"strconv"
	"strings"
)

// This is where all the conversion functions go
// This first section is for converting from a base to decimal
func hexToDecimal(hex string) int {
	hex = strings.ToUpper(hex)
	answer := 0
	exponent := float64(len(hex) - 1)

	// Since a string is essentially an array of runes, we can cycle through
	// each position in the string and parse out the number at that position
	// we then multiply that number by the appropriate power of 16
	for i := 0; i < len(hex); i++ {
		answer += fromHexMap[string(hex[i])] * int(math.Pow(16, exponent))
		exponent--
	}
	return answer
}

func octalToDecimal(octal string) int {
	answer := 0
	exponent := float64(len(octal) - 1)

	// cycling through the string to parse out each number and then multiply
	// it by the appropriate power of 8
	for i := 0; i < len(octal); i++ {
		answer += fromOctMap[string(octal[i])] * int(math.Pow(8, exponent))
		exponent--
	}
	return answer
}

func binaryToDecimal(binary string) int {
	answer := 0
	exponent := float64(len(binary) - 1)

	for i := 0; i < len(binary); i++ {
		answer += fromBinMap[string(binary[i])] * int(math.Pow(2, exponent))
		exponent--
	}
	return answer
}

// This section is for converting from decimal to another base
func decimalToHex(decimal int) string {
	answer := ""

	for decimal > 0 {
		// before dividing we want the remainder to be converted to hexadecimal
		answer += toHexMap[decimal%16]
		decimal /= 16
	}
	// before returning the answer we must reverse the string, as it is currently
	// in reverse order: LSD -> MSD, we want MSD -> LSD
	answer = reverseString(answer)
	return answer
}

func decimalToOctal(decimal int) string {
	answer := ""
	for decimal > 0 {
		// before dividing we want the remainder to be converted to octal
		answer += toOctMap[decimal%8]
		decimal /= 8
	}
	// before returning the answer we must reverse the string, as it is currently
	// in reverse order: LSD -> MSD, we want MSD -> LSD
	answer = reverseString(answer)
	return answer
}

func decimalToBinary(decimal int) string {
	answer := ""
	for decimal > 0 {
		// before dividing we want the remainder to be converted to binary
		answer += toBinMap[decimal%2]
		decimal /= 2
	}
	// before returning the answer we must reverse the string, as it is currently
	// in reverse order: LSD -> MSD, we want MSD -> LSD
	answer = reverseString(answer)
	return answer
}

// This last section is for handling all other cases
func BaseConversion(fomBase, toBase int, number string) string {
	var firstConversion int
	var finalConversion string

	switch fomBase {
	case 1:
		firstConversion = binaryToDecimal(number)
	case 2:
		firstConversion = octalToDecimal(number)
	case 3:
		firstConversion = hexToDecimal(number)
	case 4:
		// This is for when the user is converting from decimal
		firstConversion, _ = strconv.Atoi(number)
	}

	switch toBase {
	case 1:
		finalConversion = decimalToBinary(firstConversion)
	case 2:
		finalConversion = decimalToOctal(firstConversion)
	case 3:
		finalConversion = decimalToHex(firstConversion)
	case 4:
		// This is for when the user is converting to decimal
		finalConversion = strconv.Itoa(firstConversion)
	}

	return finalConversion
}

func isValidBase(fromBase int, number string) bool {
	validBase := true
	for i := 0; i < len(number); i++ {
		switch fromBase {
		case 1:
			if _, exists := fromBinMap[string(number[i])]; !exists {
				validBase = false
			}
		case 2:
			if _, exists := fromOctMap[string(number[i])]; !exists {
				validBase = false
			}
		case 3:
			if _, exists := fromHexMap[string(number[i])]; !exists {
				validBase = false
			}
		case 4:
			if _, err := strconv.Atoi(string(number[i])); err != nil {
				validBase = false
			}
		}
	}

	return validBase
}

func reverseString(number string) string {
	runeArray := []rune(number)
	for i, j := 0, len(runeArray)-1; i < j; i, j = i+1, j-1 {
		runeArray[i], runeArray[j] = runeArray[j], runeArray[i]
	}
	return string(runeArray)
}
