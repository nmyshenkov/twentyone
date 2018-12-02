package show

import "fmt"

//Card - print one player's card
func Card(card string) {
	fmt.Print("┏---┓\n|" + card + "|\n┗---┛\n")
}

//PrintDelimetr - print delimetr
func PrintDelimetr() {
	fmt.Print("▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒\n")
}

//ColorPrint - beautiful text out
func ColorPrint(text, attr string) {
	fmt.Print("\x1b[" + attr + "m" + text + "\x1b[0m")
}

//GetColorText - return color text
func GetColorText(text, attr string) string {
	return "\x1b[" + attr + "m" + text + "\x1b[0m"
}

//Hand - print player's hand
func Hand(hand []string) {
	var firstStr string
	var secondStr string
	var thirdStr string

	for _, h := range hand {
		firstStr += "┏---┓ "
		secondStr += "|" + h + "| "
		thirdStr += "┗---┛ "
	}

	firstStr += "\n"
	secondStr += "\n"
	thirdStr += "\r\n"

	fmt.Print(firstStr + secondStr + thirdStr)
}

//▒▒▒▒
//▒4♥▒
//▒▒▒▒
//

//
//  ┏----┓
//  ┃⑩  ┃
//  ┃  ♥ ┃
//  ┗----┛
//
