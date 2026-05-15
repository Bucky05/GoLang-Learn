package main

import "fmt"

var x int = 0

func main() {

	cards := newDeck()
	cards.print()
	cards.shuffle()
	fmt.Println("-------------------")
	cards.print()

}
