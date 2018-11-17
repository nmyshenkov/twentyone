package main

import (
	"fmt"
	"strconv"
	"twentyone/show"
)

//Cards - struct with information about cards
type Cards struct {
	CardValues []string
	Suits      []string
	Amounts    map[string]int
	Deck       map[string]int
}

//Player - struct with information about players
type Player struct {
	Name   string
	Hand   []string
	Point  int
	isLose bool
}

//Game - struct for main game
type Game struct {
	Pack        Cards
	Players     []Player
	PlayersLose int
	Turn        int
	isEnded     bool
}

//AddPlayer - add new player to game
func (g *Game) AddPlayer(Name string) {
	g.Players = append(g.Players, Player{
		Name: Name,
	})
}

//Start - init values and create deck
func (g *Game) Start() {
	g.Pack.Init()
	g.Pack.CreateDeck()
}

//Init - inizialization values for deck
func (c *Cards) Init() {
	c.CardValues = []string{"⑥", "⑦", "⑧", "⑨", "⑩", "В", "Д", "К", "Т"}
	c.Suits = []string{"♣", "♠", "♥", "♦"}
	c.Amounts = map[string]int{"⑥": 6, "⑦": 7, "⑧": 8, "⑨": 9, "⑩": 10, "В": 2, "Д": 3, "К": 4, "Т": 11}
}

//CreateDeck - create new deck
func (c *Cards) CreateDeck() {

	c.Deck = map[string]int{}

	for _, value := range c.CardValues {
		for _, suit := range c.Suits {
			c.Deck[value+" "+suit] = c.Amounts[value]
		}
	}
}

//GetCard - method to get card and remove this card from deck
func (c *Cards) GetCard() (string, int) {

	var key string
	var value int

	for k, v := range c.Deck {
		key = k
		value = v
		break
	}
	delete(c.Deck, key)

	return key, value
}

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

func main() {

	//TODO:: сделать добавление игроков к игре. функция main должна заключатся к инициализации, добавлении игроков и функции startGame
	fmt.Print("\nWelcome to \"21\" Game!!!\n\n")
	fmt.Print("How many players will be play?\n")

	var playerCount int

	fmt.Scan(&playerCount)

	game := Game{}

	for i := 0; i < playerCount; i++ {
		game.AddPlayer("Player" + strconv.Itoa(i))
	}
	game.Start()
	game.Turn = 1

	// fmt.Printf("%+v\n", game.Players)
	// os.Exit(1)

	for {

		if game.isEnded == true {
			fmt.Println("GAME OVER!!")
			break
		}

		//fmt.Printf("%+v\n", game.Pack.Deck)

		fmt.Println("Turn: " + strconv.Itoa(game.Turn))

		for i, player := range game.Players {
			fmt.Println(player.Name + " points: " + strconv.Itoa(player.Point) + ". Get one more card?")

			if getCommand() == false {
				break
			}

			var card string
			var point int

			card, point = game.Pack.GetCard()
			game.Players[i].Point, game.Players[i].Hand = player.Point+point, append(player.Hand, card)

			// fmt.Println(player.Name + " got: ")
			show.Card(card)
			// fmt.Println(player.Name + "'s hand: ")
			// show.Hand(player.Hand)

			fmt.Printf("%+v\n", player)

			if game.Players[i].Point > 21 {
				fmt.Println(player.Name + " lose!!")
				game.Players[i].isLose = true
				game.PlayersLose++
				break
			}

			show.PrintDelimetr()
		}

		if playerCount == game.PlayersLose {
			game.isEnded = true
		}

		game.Turn++
	}

	//fmt.Println("Your points: " + strconv.Itoa(Player1.Point))
	//fmt.Println("GAME OVER!!")
}
