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
	Pack         Cards
	Players      []Player
	PlayersCount int
	PlayersLose  int
	PlayersEnd   int
	Turn         int
	isEnded      bool
	isWon        Player
}

//AddPlayer - add new player to game
func (g *Game) AddPlayer(Name string) {
	g.Players = append(g.Players, Player{
		Name: Name,
	})
	g.PlayersCount++
}

//GetEndPlayers - get slice with end game players
func (g *Game) GetEndPlayers() []Player {
	var players []Player

	for _, p := range g.Players {
		if p.isEnd == true {
			players = append(players, p)
		}
	}

	return players
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
