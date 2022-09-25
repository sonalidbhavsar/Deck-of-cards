// deck package implements operations on card's deck.
// This package can be used to build cards games.
package deck

import (
	"math/rand"
	"sort"
)

// returns sorted deck of cards
func GetSingleDeck() []Card {
	var cardsDeck []Card
	var s Suit
	var c CardNumber
	for s = 0; s < NumberOfSuitTypes; s++ {
		for c = 0; c < NumberOfCardTypes; c++ {
			card := new(Card)
			card.SuitType = s
			card.CardType = c
			cardsDeck = append(cardsDeck, *card)
		}
	}
	return cardsDeck
}

// returned n decks, all in sorted order
func New(deckCount int) []Card {
	var cards []Card
	for i := 0; i < deckCount; i++ {
		c := GetSingleDeck()
		cards = append(cards, c...)
	}
	return cards
}

// Shuffles the deck
func Shuffle(cards []Card) {
	maxLen := len(cards) - 1
	for i := 0; i <= maxLen/2; i++ {
		j, k := rand.Intn(maxLen), rand.Intn(maxLen)
		cards[j], cards[k] = cards[k], cards[j]
	}
}

// sorts the deck.
// TODO: sort for multi deck
func SortCards(cards []Card) []Card {
	c := CardsSorter{cards: cards}
	sort.Sort(c)
	return c.cards
}

// given two cards, determines if first card should come before the second in sorted list
func isALesserThanB(a Card, b Card) bool {
	if a.SuitType == b.SuitType {
		return a.CardType < b.CardType
	}
	return a.SuitType < b.SuitType
}

// filters the cards with given cardNumber and suitType from the given deck
func FilterCards(cards []Card, number string, suit string) {
	lastIndex := len(cards)
	for i := 0; i < lastIndex; {
		card := cards[i]
		if (number == "" || card.CardType.String() == number) && (suit == "" || card.SuitType.String() == suit) {
			for j := i; j < lastIndex-1; j++ {
				cards[j] = cards[j+1]
			}
			lastIndex--
		} else {
			i++
		}
	}
	cards = cards[:lastIndex]
}

// adds n jokers at the end of the deck
func AddJokers(cards []Card, count int) []Card {
	for i := 0; i < count; i++ {
		cards = append(cards, Card{
			SuitType: -1,
			CardType: Joker,
		})
	}
	return cards
}
