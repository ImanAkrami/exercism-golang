package partyrobot

import (
	"fmt"
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {

	return "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	msg := ""
	msg += "Welcome to my party, "
	msg += name
	msg += "!\n"
	msg += "You have been assigned to table "
	var tableNumStr string
	if table < 10 {
		tableNumStr = "00" + strconv.Itoa(table)
	} else if table < 100 {
		tableNumStr = "0" + strconv.Itoa(table)

	} else {
		tableNumStr = strconv.Itoa(table)
	}
	msg += tableNumStr
	msg += ". Your table is "
	msg += direction
	msg += ", exactly "
	dist1Decimal := fmt.Sprintf("%.1f", distance)
	msg += dist1Decimal
	msg += " meters from here."
	msg += "\n"
	msg += "You will be sitting next to "
	msg += neighbor
	msg += "."
	return msg
}

// Welcome to my party, Christiane!
// You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
// You will be sitting next to Frank.
