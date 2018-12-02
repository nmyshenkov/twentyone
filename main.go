package main

import (
	"fmt"
	"strconv"
	"twentyone/show"
)

//getCommand - get command
func getCommand() bool {
	var command string

	for {

		fmt.Scan(&command)

		switch command {
		case "y":
			return true
		case "n":
			return false
		default:
			fmt.Println("Wrong command. Press 'y' or 'n")
		}
	}
}

//isEndGame - check ending game
func isEndGame(g Game) bool {
	if g.isEnded {
		return true
	}

	if g.PlayersCount == g.PlayersLose {
		show.ColorPrint("All players lost!! Game over\n", "31")
		return true
	}

	if g.PlayersCount == g.PlayersEnd+g.PlayersLose {
		playerEnds := g.GetEndPlayers()
		for _, p := range playerEnds {
			if p.Point > g.isWon.Point {
				g.isWon = p
			}
		}
		show.ColorPrint("Player "+g.isWon.Name+" WON with "+strconv.Itoa(g.isWon.Point)+" points!!!\n", "32")
		return true
	}

	return false
}

func main() {

	show.ColorPrint("\nWelcome to \"21\" Game!!!\n\n", "1;33")
	show.ColorPrint("How many players will be play? ", "4")

	var playerCount int

	fmt.Scan(&playerCount)

	game := Game{}

	for i := 0; i < playerCount; i++ {
		game.AddPlayer("Player" + strconv.Itoa(i+1))
	}
	game.Start()
	game.Turn = 1

	for {

		if isEndGame(game) == true {
			game.isEnded = true
		}

		if game.isEnded == true {
			fmt.Println("GAME OVER!!")
			break
		}

		fmt.Println("Turn: " + strconv.Itoa(game.Turn))

		for i, player := range game.Players {

			if player.isLose == true || player.isEnd == true {
				continue
			}

			show.PrintDelimetr()

			fmt.Println(player.Name + " points: " + strconv.Itoa(player.Point) + ". Get one more card?")

			if getCommand() == false {
				game.Players[i].isEnd = true
				game.PlayersEnd++
				continue
			}

			var card string
			var point int

			card, point = game.Pack.GetCard()
			game.Players[i].Point, game.Players[i].Hand = player.Point+point, append(player.Hand, card)

			fmt.Println(player.Name + " got: ")
			show.Card(card)
			fmt.Println(player.Name + "'s hand: ")
			show.Hand(game.Players[i].Hand)

			if game.Players[i].Point > 21 {
				show.ColorPrint(player.Name+" lose with "+strconv.Itoa(game.Players[i].Point)+" points!!!\n", "31")
				game.Players[i].isLose = true
				game.PlayersLose++
				continue
			}
		}

		if playerCount == game.PlayersLose {
			game.isEnded = true
		}

		game.Turn++
	}
}
