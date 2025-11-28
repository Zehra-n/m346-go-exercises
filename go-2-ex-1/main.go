package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   byte
	Month byte
	Year  uint16
}

type Profile struct {
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		FullName: FullName{ 
			FirstName: "Zehra",
			LastName:  "Nuhiu",
		},
		BirthDate: BirthDate{
			Day:   26,
			Month: 6,
			Year:  2007,
		},
		NumberOfSiblings: 2,   // TODO: adjust
		ZodiacSign:       '\u264B', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
