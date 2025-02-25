package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var choice string

	ClearScreen()
	fmt.Println("Welcome to the HOB Converter!")
	fmt.Println()
	fmt.Println("This simple CLI tool will allow you to convert numbers from\n" +
		"one base to another. The base systems that you can convert between are:\n" +
		"Binary, Octal, Hexadecimal, and Decimal.")
	fmt.Println()

	// This is the main loop of the program. It will keep asking the user if they
	// want to convert a number until they choose to quit
	for {
		fmt.Println("Would you like to convert a number? [y/n]")
		fmt.Scan(&choice)
		choice = strings.ToLower(choice)
		switch choice {
		case "y":
			ConvertNumberBase()
		case "n":
			fmt.Println("Goodbye")
			return
		}
	}
}

func subMenu() {
	fmt.Println("1) Binary")
	fmt.Println("2) Octal")
	fmt.Println("3) Hexadecimal")
	fmt.Println("4) Decimal")
}

func showConversion(fromBase, toBase int) {
	var fromWord, toWord string

	switch fromBase {
	case 1:
		fromWord = "Binary"
	case 2:
		fromWord = "Octal"
	case 3:
		fromWord = "Hexadecimal"
	case 4:
		fromWord = "Decimal"
	default:
		fromWord = ""
	}

	switch toBase {
	case 1:
		toWord = "Binary"
	case 2:
		toWord = "Octal"
	case 3:
		toWord = "Hexadecimal"
	case 4:
		toWord = "Decimal"
	default:
		toWord = ""
	}

	fmt.Printf("%s => %s\n", fromWord, toWord)
}

func ConvertNumberBase() {
	var fromBase, toBase int
	var number string
	validBase := false

	ClearScreen()
	subMenu()
	fmt.Println()
	for {
		fmt.Println("Choose a base to convert from: ")
		fmt.Scan(&fromBase)

		if fromBase < 1 || fromBase > 4 {
			fmt.Println("Not a valid base")
			fmt.Println("Please try again")
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	ClearScreen()
	showConversion(fromBase, toBase)
	subMenu()
	fmt.Println()
	for {
		fmt.Println("Choose a base to convert to: ")
		fmt.Scan(&toBase)

		if toBase < 1 || toBase > 4 {
			fmt.Println("Not a valid base")
			fmt.Println("Please try again")
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	for !validBase {
		ClearScreen()
		showConversion(fromBase, toBase)
		fmt.Println("Enter a number: ")
		fmt.Scan(&number)
		validBase = isValidBase(fromBase, number)
		if !validBase {
			fmt.Println("Not a valid number for the chosen base")
			fmt.Println("Please try again")
			// add a bit of delay so user can process the message.
			time.Sleep(2 * time.Second)
		}
	}

	fmt.Println("The converted number is: ")
	fmt.Println(ConvertBase(fromBase, toBase, number))
}
