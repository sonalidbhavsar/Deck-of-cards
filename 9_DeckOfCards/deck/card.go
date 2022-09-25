package deck

// represents a single card
type Card struct {
	SuitType Suit
	CardType CardNumber
}

// sorter used for sorting of deck.
type CardsSorter struct {
	cards []Card
}

// determines if card A should come before card B in sorted deck
func (c CardsSorter) Less(i, j int) bool {
	return isALesserThanB(c.cards[i], c.cards[j])
}

// swaps cards from position i and j in the deck
func (c CardsSorter) Swap(i, j int) {
	c.cards[i], c.cards[j] = c.cards[j], c.cards[i]
}

// returns length of deck
func (c CardsSorter) Len() int {
	return len(c.cards)
}
