package main

import ("fmt"
 "log")
import "rsc.io/quote/v4"
import "example.com/greetings"


func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err:= greetings.Greet("")

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(message + quote.Go())
}