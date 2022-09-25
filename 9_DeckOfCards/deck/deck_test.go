package deck

import (
	"testing"
)

func TestNew(t *testing.T) {
	cards := New(1)
	if len(cards) != NumberOfCardTypes*NumberOfSuitTypes {
		t.Error()
	}
	cards = New(2)
	if len(cards) != 2*13*4 {
		t.Error()
	}
}

func TestShuffle(t *testing.T) {
	cards := New(1)
	Shuffle(cards)
	if len(cards) != NumberOfCardTypes*NumberOfSuitTypes {
		t.Error()
	}
	cards = New(2)
	if len(cards) != 2*NumberOfCardTypes*NumberOfSuitTypes {
		t.Error()
	}
}

func TestFilterCards(t *testing.T) {
	cards := New(1)
	number := _J.String()
	suit := Spade.String()
	FilterCards(cards, number, suit)
	for _, card := range cards {
		if (number == "" || card.CardType.String() == number) && (suit == "" || card.SuitType.String() == suit) {
			t.Error("Error: Filter cards, cardType: " + card.CardType.String() + " suitType: " + card.SuitType.String())
		}
	}
}

func TestSortCards(t *testing.T) {
	cards := New(1)
	Shuffle(cards)
	SortCards(cards)
	if len(cards) != 13*4 {
		t.Error("Error: Incorrect length")
	}
	if cards[0].SuitType != Spade || cards[0].CardType != _A {
		t.Error("Error: First card in sorted deck is not Spade A")
	}
	if cards[13].SuitType != Diamond || cards[13].CardType != _A {
		t.Error("Error: First card of second suit in sorted deck is not Spade A")
	}
	if cards[14].SuitType != Diamond || cards[14].CardType != _2 {
		t.Error("Error: Second card of second suit in sorted deck is not Spade A")
	}

}
