package deck

//go:generate stringer -type=Suit
//go:generate stringer -type=CardNumber

type Suit int
type CardNumber int

// types of suits
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
)

// types of card in each suit
const (
	_A CardNumber = iota
	_2
	_3
	_4
	_5
	_6
	_7
	_8
	_9
	_10
	_J
	_Q
	_K
	Joker
)
