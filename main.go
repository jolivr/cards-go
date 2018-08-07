package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	hand := deck{}
	remainingDeck := cards
	cardCount := len(cards)
	handCount := 0
	handSize := 5

	for cardCount > 0 {
		if cardCount < handSize {
			handSize = cardCount
		}

		hand, remainingDeck = deal(remainingDeck, handSize)

		cardCount = len(remainingDeck)
		handCount++

		message := fmt.Sprintf("Hand %d", handCount)
		fmt.Println(message)
		hand.print()
	}

}
