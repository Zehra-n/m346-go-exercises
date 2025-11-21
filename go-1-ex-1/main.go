package main

import "fmt"

func main() {
	firstName := "Zehra"
	lastName := "Nuhiu"
	dayOfBirth := 26
	monthOfBirth := 6
	yearOfBirth := 2007
	numberOfSiblings := 2
	heightInMeters := 1.68
	zodiacSign := '\u264B'
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
