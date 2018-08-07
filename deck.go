package main

import (
	"fmt"
)

//Creating new type of 'deck'
//Slice of string

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
