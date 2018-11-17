package show

import "fmt"

func Card(card string) {
	fmt.Print("┏---┓\n|" + card + "|\n┗---┛\n")
}

func PrintDelimetr() {
	fmt.Print("-------------------------\n")
}

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

//  ┏----┓
//  ┃⑩  ┃
//  ┃  ♥ ┃
//  ┗----┛
//
//
//
