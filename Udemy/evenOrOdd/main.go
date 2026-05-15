package main

import "fmt"

func main() {
	intArr := [11]int8{}

	for i := range 11 {
		intArr[i] = int8(i)
	}

	for _, element := range intArr {
		if element%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
