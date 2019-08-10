package main

import "fmt"

// Create a new type of deck
//  which is a slicce of strings

type deck []string //string typeのsubclass 的な

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
