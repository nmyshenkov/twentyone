package main

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
	isEnd  bool
}

//Game - struct for main game
type Game struct {
	Pack          Cards
	PlayerIDs     []int
	Players       map[int]*Player
	PlayersInGame []int
	PlayersCount  int
	PlayersLose   int
	PlayersEnd    int
	Turn          int
	isEnded       bool
	isWon         Player
}

//Init - init struct
func (g *Game) Init() {
	g.Players = make(map[int]*Player)
}

//AddPlayer - add new player to game
func (g *Game) AddPlayer(Name string) {

	id := len(g.Players) + 1

	g.Players[id] = &Player{
		Name: Name,
	}
	g.PlayerIDs = append(g.PlayerIDs, id)
	g.PlayersInGame = append(g.PlayersInGame, id)
	g.PlayersCount++
}

//LosePlayer - mark who lose in the game
func (g *Game) LosePlayer(id int) {
	for i, el := range g.PlayersInGame {
		if el == id {
			g.PlayersInGame = append(g.PlayersInGame[:i], g.PlayersInGame[i+1:]...)
			break
		}
	}
}

//GetEndPlayers - get slice with end game players
func (g *Game) GetEndPlayers() map[int]*Player {
	players := make(map[int]*Player)

	for i, p := range g.Players {
		if p.isEnd == true {
			players[i] = p
		}
	}

	return players
}

//Start - init values and create deck
func (g *Game) Start() {
	g.Pack.Init()
	g.Pack.CreateDeck()

	var startCards = 2
	var card string
	var point int

	for i := 0; i < startCards; i++ {
		for i, player := range g.Players {
			card, point = g.Pack.GetCard()
			g.Players[i].Point, g.Players[i].Hand = player.Point+point, append(player.Hand, card)
		}
	}
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
